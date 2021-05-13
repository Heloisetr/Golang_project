package http

import (
	"ads/domain"
	"ads/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type updateAdRequest struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Price       float32        `json:"price"`
	Picture     domain.Picture `json:"picture"`
}

func UpdateAdHandler(cmd usecase.UpdateAdCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		adID := c.Param("ad_id")
		if adID == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		updateAdReq := &updateAdRequest{}
		err := c.BindJSON(updateAdReq)
		if err != nil {
			logrus.WithError(err).Error("error while binding JSON")
			c.Status(http.StatusBadRequest)
			return
		}

		ad, err := cmd(c.Request.Context(), adID, domain.Ad{
			Title:       updateAdReq.Title,
			Description: updateAdReq.Description,
			Price:       updateAdReq.Price,
			Picture:     updateAdReq.Picture,
		})
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
