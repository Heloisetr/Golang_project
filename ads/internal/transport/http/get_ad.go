package http

import (
	"ads/domain"
	"ads/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdHandler(cmd usecase.GetAdCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		adID := c.Param("ad_id")
		if adID == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		ad, err := cmd(c.Request.Context(), adID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAdNotFound):
				c.Status(http.StatusNotFound)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusOK, ad)
	}
}
