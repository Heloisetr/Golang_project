package http

import (
	"ads/domain"
	"ads/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteAdHandler(cmd usecase.DeleteAdCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		adID := c.Param("ad_id")
		if adID == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		err := cmd(c.Request.Context(), c.GetString("token"), adID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrTokenParsing):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrTokenParsing.Error()})
			case errors.Is(err, domain.ErrCantDelete):
				c.JSON(http.StatusConflict, gin.H{"message": domain.ErrCantDelete})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusOK, "Ad deleted")
	}
}
