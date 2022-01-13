package http

import (
	"account/domain"
	"account/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getAccountResponse struct {
	AccountID string  `json:"account_id"`
	Email     string  `json:"email"`
	Login     string  `json:"login"`
	Balance   float32 `json:"balance"`
}

func GetAccountHandler(cmd usecase.GetAccountCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		rqAccountID := c.Param("account_id")

		if rqAccountID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "No Account ID provide"})
			return
		}

		account, err := cmd(c.Request.Context(), c.GetString("token"), rqAccountID)

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAccountNotFound):
				c.Status(http.StatusNotFound)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}

		c.JSON(http.StatusOK, getAccountResponse{
			AccountID: account.AccountID,
			Email:     account.Email,
			Login:     account.Login,
			Balance:   account.Balance,
		})
	}
}
