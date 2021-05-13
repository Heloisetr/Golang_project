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

		ads, err := cmd(c.Request.Context(), userID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAdNotFound):
				c.Status(http.StatusNotFound)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusOK, ads)
	}
}
