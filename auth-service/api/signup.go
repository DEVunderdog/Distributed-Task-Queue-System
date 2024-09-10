package api

import (
	"net/http"
	"time"

	database "github.com/DEVunderdog/auth-service/database/sqlc"
	"github.com/DEVunderdog/auth-service/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type signupUserRequest struct {
	Email           string `json:"email" binding:"required,email,domain_email"`
	Password        string `json:"password" binding:"required,strong_password"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type userResponse struct {
	Id             int64     `json:"id"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func newUserResponse(user database.User) userResponse {
	return userResponse{
		Id: user.ID,
		Email: user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}
}

func (server *Server) signupUser(ctx *gin.Context) {

	var req signupUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		if validationErros, ok := err.(validator.ValidationErrors); ok {
			errorMessage := validationErrorResponse(validationErros)
			server.enhanceHTTPResponse(ctx, http.StatusBadRequest, "Error in user request body", errorMessage)
			return
		}
		server.enhanceHTTPResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	exists , err := server.store.CheckForExistingUser(ctx, req.Email)
	if exists {
		server.enhanceHTTPResponse(ctx, http.StatusConflict, "user already exists", nil)
		return
	}

	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if req.Password != req.ConfirmPassword {
		server.enhanceHTTPResponse(ctx, http.StatusBadRequest, "Password does not match", nil)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	arg := database.CreateUserParams{
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	response := newUserResponse(user)
	server.enhanceHTTPResponse(ctx, http.StatusCreated, "user created successfully", response)
}
