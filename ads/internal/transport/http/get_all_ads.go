package http

import (
	"ads/domain"
	"ads/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllsAdHandler(cmd usecase.GetAllAdCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")
		if userID == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		ads, err := cmd(c.Request.Context(), c.GetString("token"), userID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAdNotFound):
				c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrAdNotFound.Error()})
			case errors.Is(err, domain.ErrUnauthorized):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrUnauthorized.Error()})
			case errors.Is(err, domain.ErrTokenParsing):
				c.JSON(http.StatusBadRequest, gin.H{"message": domain.ErrTokenParsing.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusOK, ads)
	}
}
