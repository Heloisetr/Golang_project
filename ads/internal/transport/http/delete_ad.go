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

		err := cmd(c.Request.Context(), adID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAdNotFound):
				c.Status(http.StatusNotFound)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusOK, "Ad erased")
	}
}
