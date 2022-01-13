package http

import (
	"account/domain"
	"account/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(cmd usecase.LoginCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		loginReq := &loginRequest{}
		err := c.BindJSON(loginReq)

		if err != nil {
			logrus.WithError(err).Error("error while binding json body")
			c.Status(http.StatusBadRequest)
			return
		}

		token, err := cmd(c.Request.Context(), domain.Login{
			Email:    loginReq.Email,
			Password: loginReq.Password,
		})

		if err != nil {
			switch {
			case errors.Is(err, domain.ErrEmailNotFound):
				c.JSON(http.StatusBadRequest, gin.H{"message": domain.ErrEmailNotFound.Error()})
			case errors.Is(err, domain.ErrWrongPassword):
				c.JSON(http.StatusForbidden, gin.H{"message": domain.ErrWrongPassword.Error()})
			case errors.Is(err, domain.ErrTokenCreation):
				c.JSON(http.StatusConflict, gin.H{"message": domain.ErrTokenCreation.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			}
			return
		}

		c.JSON(http.StatusOK, loginResponse{Token: token})
	}
}
