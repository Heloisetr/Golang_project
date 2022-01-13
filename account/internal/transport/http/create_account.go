package http

import (
	"account/domain"
	"account/internal/usecase"
	"account/internal/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type createAccountRequest struct {
	Email    string  `json:"email"`
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Balance  float32 `json:"balance"`
}

type createAccountResponse struct {
	AccountID string `json:"account_id"`
}

func CreateAccountHandler(cmd usecase.CreateAccountCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		createAccountReq := &createAccountRequest{}
		err := c.BindJSON(createAccountReq)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error while binding json body"})
			return
		}

		if utils.CheckAccount(domain.Account{
			Email:    createAccountReq.Email,
			Login:    createAccountReq.Login,
			Password: createAccountReq.Password,
			Balance:  createAccountReq.Balance,
		}) == false {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error while binding json body"})
			return
		}

		hashPassword, err := bcrypt.GenerateFromPassword([]byte(createAccountReq.Password), bcrypt.MinCost)

		accountID, err := cmd(c.Request.Context(), domain.Account{
			Email:    createAccountReq.Email,
			Login:    createAccountReq.Login,
			Password: string(hashPassword),
			Balance:  createAccountReq.Balance,
		})

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrEmailAlreadyUsed):
				c.JSON(http.StatusConflict, gin.H{"message": domain.ErrEmailAlreadyUsed.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}

		c.JSON(http.StatusCreated, createAccountResponse{AccountID: accountID})
	}
}
