package api

import (
	"net/http"

	database "github.com/DEVunderdog/auth-service/database/sqlc"
	"github.com/DEVunderdog/auth-service/token"
	"github.com/DEVunderdog/auth-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgtype"
)

type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email,domain_email"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			errorMessage := validationErrorResponse(validationErrors)
			server.enhanceHTTPResponse(ctx, http.StatusBadRequest, "Error in user request body", errorMessage)
			return
		}
		server.enhanceHTTPResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := server.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	
	correctPassword, err := utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	if !correctPassword {
		server.enhanceHTTPResponse(ctx, http.StatusBadRequest, "your given password is incorrect", nil)
		return
	}

	activeKey,err := token.GetActiveJWTKey(ctx, true, server.store)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if len(activeKey) <= 0 {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, "No keys available, generate one", nil)
		return
	}

	key := activeKey[0]

	pvtKey, err := token.GetPrivateKey(key.PrivateKey, []byte(server.config.Passphrase))
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	pubKey, err := token.GetPublicKey(key.PublicKey)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	signedToken, signedRefreshToken, err := token.GenerateToken(pvtKey, uint(user.ID), user.Email, server.config)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	tokenExpirationTime, err := token.GetExpirationTime(signedToken, pubKey)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	refreshTokenExpiration, err := token.GetExpirationTime(signedRefreshToken, pubKey)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	arg := database.CreateSessionParams{
		UserID: user.ID,
		Token: signedToken,
		RefreshToken: signedRefreshToken,
		TokenExpiresAt: pgtype.Timestamptz{
			Time: *tokenExpirationTime,
			Valid: true,
		},
		RefreshTokenExpiresAt: pgtype.Timestamptz{
			Time: *refreshTokenExpiration,
			Valid: true,
		},
		IsActive: pgtype.Bool{
			Bool: true,
			Valid: true,
		},
		Ip: pgtype.Text{
			String: ctx.ClientIP(),
			Valid: true,
		},
		UserAgent: pgtype.Text{
			String: ctx.Request.UserAgent(),
			Valid: true,
		},
	}

	_, err = server.store.CreateSession(ctx, arg)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response := loginUserResponse{
		AccessToken: signedToken,
		RefreshToken: signedRefreshToken,
	}

	server.enhanceHTTPResponse(ctx, http.StatusOK, "User logged in successfully", response)
}
