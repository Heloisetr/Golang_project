package http

import (
	"account/domain"
	"account/external"
	"account/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type deleteAccountResponse struct {
	Message string `json:"message"`
}

func DeleteAccountHandler(cmd usecase.DeleteAccountCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		rqAccountID := c.Param("account_id")

		if rqAccountID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No Account ID provide"})
			return
		}

		message, err := cmd(c.Request.Context(), c.GetString("token"), rqAccountID)

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAccountNotFound):
				c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrAccountNotFound.Error()})
			case errors.Is(err, domain.ErrAccessUnauthorized):
				c.JSON(http.StatusUnauthorized, gin.H{"message": domain.ErrAccessUnauthorized.Error()})
			case errors.Is(err, external.ErrExternalGetAds):
				c.JSON(http.StatusConflict, gin.H{"message": external.ErrExternalGetAds.Error()})
			case errors.Is(err, external.ErrExternalDeleteAd):
				c.JSON(http.StatusConflict, gin.H{"message": external.ErrExternalDeleteAd.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}

		c.JSON(http.StatusOK, deleteAccountResponse{
			Message: message,
		})
	}
}
