package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kweku-xvi/todo/api/v1/models"
	"github.com/kweku-xvi/todo/internal/config"
	"github.com/kweku-xvi/todo/internal/database"
)

func CheckAuth(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")

	if tokenStr == "" {
		c.JSON(401, gin.H{
			"error": "authorization header is missing",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ENV.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{
			"error": "invalid or expired token",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(401, gin.H{
			"error": "invalid token",
		})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(401, gin.H{
			"error": "token expired",
		})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	database.DB.Where("id = ?", claims["id"]).Find(&user)

	if user.ID == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("currentUser", user)
	c.Next()
}
