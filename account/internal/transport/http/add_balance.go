package http

import (
	"account/domain"
	"account/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type addBalanceRequest struct {
	Balance float32 `json:"balance"`
}

type addBalanceResponse struct {
	Message string `json:"message"`
}

func AddBalanceHandler(cmd usecase.AddBalanceCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		addBalanceReq := &addBalanceRequest{}
		err := c.BindJSON(addBalanceReq)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error while binding json body"})
			return
		}

		message, err := cmd(c.Request.Context(), c.GetString("token"), addBalanceReq.Balance)

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrAccountNotFound):
				c.JSON(http.StatusNotFound, gin.H{"message": domain.ErrAccountNotFound.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}

		c.JSON(http.StatusOK, addBalanceResponse{Message: message})
	}
}
