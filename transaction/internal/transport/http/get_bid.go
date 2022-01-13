package http

import (
	"errors"
	"net/http"
	"transaction/domain"
	"transaction/internal/usecase"

	"github.com/gin-gonic/gin"
)

func GetBidHandler(cmd usecase.GetBidCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		bidID := c.Param("bid_id")
		if bidID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad parameter provide"})
			return
		}

		bid, err := cmd(c.Request.Context(), bidID, c.GetString("token"))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrBidNotFound):
				c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrAdNotFound.Error()})
			case errors.Is(err, domain.ErrToken):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrToken.Error()})
			case errors.Is(err, domain.ErrUnauthorized):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrUnauthorized.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusOK, bid)
	}
}
