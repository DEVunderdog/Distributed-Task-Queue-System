package api

import (
	"net/http"
	"github.com/DEVunderdog/auth-service/token"
	"github.com/gin-gonic/gin"
)

func(server *Server) logoutUser (ctx *gin.Context) {

	claims, exists := ctx.Get("claims")

	if !exists {
		server.enhanceHTTPResponse(ctx, http.StatusUnauthorized, "user not authorized", nil)
		return
	}

	userClaims, ok := claims.(*token.Claims)
	if !ok {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, "invalid claims", nil)
		return
	}

	user, err := server.store.GetUserByID(ctx, int64(userClaims.UserID))
	if err != nil {
		server.enhanceHTTPResponse(ctx, http.StatusInternalServerError, "coudn't find the user", nil)
		return
	}

	updatedSession, err := server.store.LoggedOutSession(ctx, user.ID)

	server.enhanceHTTPResponse(ctx, http.StatusOK, "user logged out successfully", updatedSession)

}