package http

import (
	"errors"
	"net/http"
	"transaction/domain"
	"transaction/external"
	"transaction/internal/infrastructure/types"
	"transaction/internal/infrastructure/utils"
	"transaction/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type createBidRequest struct {
	AdId     string  `json:"ad_id"`
	BidPrice float32 `json:"bid_price"`
	Message  string  `json:"message"`
	Status   string  `json:"status"`
}

type createBidResponse struct {
	BidID string `json:"bid_id"`
}

func CreateBidHandler(cmd usecase.CreateBidCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		createBidReq := &createBidRequest{}
		err := c.BindJSON(createBidReq)
		if err != nil {
			logrus.WithError(err).Error("error while binding JSON")
			c.Status(http.StatusBadRequest)
			return
		}

		accountId, errParse := utils.ParseToken(c.GetString("token"))

		if errParse != nil {
			c.JSON(http.StatusUnauthorized, errParse.Error())
		}

		bidID, err := cmd(c.Request.Context(), domain.Bid{
			UserID:   accountId,
			AdId:     createBidReq.AdId,
			BidPrice: createBidReq.BidPrice,
			Message:  createBidReq.Message,
			Status:   createBidReq.Status,
		}, c.GetString("token"))
		if err != nil {
			switch {
			case errors.Is(err, external.ErrExternalGetAccount):
				c.JSON(http.StatusNotFound, gin.H{"message": external.ErrExternalGetAccount.Error()})
			case errors.Is(err, external.ErrExternalNotEnoughBalance):
				c.JSON(http.StatusConflict, gin.H{"message": external.ErrExternalNotEnoughBalance.Error()})
			case errors.Is(err, types.ErrTypeStatusInvalid):
				c.JSON(http.StatusBadRequest, gin.H{"message": types.ErrTypeStatusInvalid.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusCreated, createBidResponse{BidID: bidID})
	}
}
