package http

import (
	"account/domain"
	"account/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type updateAccountRequest struct {
	Email string `json:"email,omitempty"`
	Login string `json:"login,omitempty"`
}

type updateAccountResponse struct {
	AccountID string  `json:"accountid"`
	Email     string  `json:"email"`
	Login     string  `json:"login"`
	Balance   float32 `json:"balance"`
}

func UpdateAccountHandler(cmd usecase.UpdateAccountCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		updateAccountReq := &updateAccountRequest{}
		err := c.BindJSON(updateAccountReq)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error while binding json body"})
			return
		}

		account, err := cmd(c.Request.Context(), c.GetString("token"), domain.UpdateAccount{
			Email: updateAccountReq.Email,
			Login: updateAccountReq.Login,
		})

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAccountNotFound):
				c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrAccountNotFound.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}

		c.JSON(http.StatusOK, account)
	}
}
