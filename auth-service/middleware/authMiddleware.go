package middleware

import (
	"net/http"
	"strings"

	database "github.com/DEVunderdog/auth-service/database/sqlc"
	"github.com/DEVunderdog/auth-service/token"
	"github.com/DEVunderdog/auth-service/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(config utils.Config, store database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		key, err := token.GetActiveJWTKey(c , true, store)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		activeKey := key[0]

		publicKey, err := token.GetPublicKey(activeKey.PublicKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Au"})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := token.ValidateToken(tokenString, publicKey, config)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}