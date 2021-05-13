package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		tokenSplitted := strings.Split(token, "Bearer ")
		if len(tokenSplitted) != 2 {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}
		userID := tokenSplitted[1]
		c.Set("user_id", userID)
	}
}
