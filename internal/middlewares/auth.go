package middlewares

import (
	"net/http"
	"spk-be/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is requried"})
// 			c.Abort()
// 			return
// 		}

// 		parts := strings.Split(authHeader, " ")
// 		if len(parts) != 2 || parts[0] != "Bearer" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString := parts[1]

// 		userId, err := utils.ValidateJWT(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("user_id", userId)

// 		c.Next()
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		if cookieToken, err := c.Cookie("access_token"); err == nil {
			token = cookieToken
		}
		if token == "" {
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}

		userID, err := utils.ValidateJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
