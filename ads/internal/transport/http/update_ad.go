package http

import (
	"ads/domain"
	"ads/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateAdRequest struct {
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	Price       float32        `json:"price,omitempty"`
	Picture     domain.Picture `json:"picture,omitempty"`
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
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error while binding json body"})
			return
		}

		ad, err := cmd(c.Request.Context(), c.GetString("token"), adID, domain.UpdateAd{
			Title:       updateAdReq.Title,
			Description: updateAdReq.Description,
			Price:       updateAdReq.Price,
			Picture:     updateAdReq.Picture,
		})
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAdNotFound):
				c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrAdNotFound.Error()})
			case errors.Is(err, domain.ErrUnauthorized):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrUnauthorized.Error()})
			case errors.Is(err, domain.ErrTokenParsing):
				c.JSON(http.StatusBadRequest, gin.H{"message": domain.ErrTokenParsing.Error()})
			case errors.Is(err, domain.ErrCantUpdate):
				c.JSON(http.StatusConflict, gin.H{"message": domain.ErrCantUpdate.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}
		c.JSON(http.StatusOK, ad)
	}
}
