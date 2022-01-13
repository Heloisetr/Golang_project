package http

import (
	"ads/domain"
	"ads/internal/usecase"
	"ads/internal/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type createAdRequest struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       float32        `json:"price"`
	Picture     domain.Picture `json:"picture"`
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

		accountID, errParse := utils.ParseToken(c.GetString("token"))

		if errParse != nil {
			c.JSON(http.StatusUnauthorized, errParse.Error())
		}

		adID, err := cmd(c.Request.Context(), domain.Ad{
			UserID:      accountID,
			Title:       createAdReq.Title,
			Description: createAdReq.Description,
			Price:       createAdReq.Price,
			Picture:     createAdReq.Picture,
		})
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrTokenParsing):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrTokenParsing.Error()})
			case errors.Is(err, domain.ErrCreate):
				c.JSON(http.StatusBadRequest, gin.H{"message": domain.ErrCreate.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusCreated, createAdResponse{AdID: adID})
	}
}
