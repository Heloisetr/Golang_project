package http

import (
	"ads/domain"
	"ads/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type createAdRequest struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       float32        `json:"price"`
	Picture     domain.Picture `json:"picture"`
	UserID      string         `json:"user_id"`
}

type createAdResponse struct {
	AdID string `json:"ad_id"`
}

func CreateAdHandler(cmd usecase.CreateAdCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		createAdReq := &createAdRequest{}
		err := c.BindJSON(createAdReq)
		if err != nil {
			logrus.WithError(err).Error("error while binding JSON")
			c.Status(http.StatusBadRequest)
			return
		}

		adID, err := cmd(c.Request.Context(), domain.Ad{
			UserID:      createAdReq.UserID,
			Title:       createAdReq.Title,
			Description: createAdReq.Description,
			Price:       createAdReq.Price,
			Picture:     createAdReq.Picture,
		})
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusCreated, createAdResponse{AdID: adID})
	}
}
