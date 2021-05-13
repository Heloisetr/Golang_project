package http

import (
	"ads/domain"
	"ads/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetByKeysAdHandler(cmd usecase.GetByKeysAdCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		keyword := c.Param("keyword")

		if keyword == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		ads, err := cmd(c.Request.Context(), keyword)
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
