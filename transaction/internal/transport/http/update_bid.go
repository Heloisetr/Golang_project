package http

import (
	"errors"
	"net/http"
	"transaction/domain"
	"transaction/internal/infrastructure/types"
	"transaction/internal/usecase"

	"github.com/gin-gonic/gin"
)

type updateBidRequest struct {
	Status string `json:"status"`
}

func UpdateBidHandler(cmd usecase.UpdateBidCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		bidID := c.Param("bid_id")
		if bidID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad parameter provide"})
			return
		}

		updateBidReq := &updateBidRequest{}
		err := c.BindJSON(updateBidReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error while binding JSON"})
			return
		}

		bid, err := cmd(c.Request.Context(), bidID, domain.Bid{
			Status: updateBidReq.Status,
		}, c.GetString("token"))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrBidNotFound):
				c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrAdNotFound.Error()})
			case errors.Is(err, domain.ErrToken):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrToken.Error()})
			case errors.Is(err, domain.ErrCantUpdate):
				c.JSON(http.StatusConflict, gin.H{"message": domain.ErrCantUpdate.Error()})
			case errors.Is(err, domain.ErrUnauthorized):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrUnauthorized.Error()})
			case errors.Is(err, types.ErrTypeStatusInvalid):
				c.JSON(http.StatusBadRequest, gin.H{"message": types.ErrTypeStatusInvalid.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusOK, bid)
	}
}
