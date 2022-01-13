package http

import (
	"epitech/deliveats/user/domain"
	"epitech/deliveats/user/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type createUserRequest struct {
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Address  domain.Address `json:"address"`
}

type createUserResponse struct {
	UserID string `json:"user_id"`
}

func CreateUserHandler(cmd usecase.CreateUserCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		createUserReq := &createUserRequest{}
		err := c.BindJSON(createUserReq)
		if err != nil {
			logrus.WithError(err).Error("error while binding JSON")
			c.Status(http.StatusBadRequest)
			return
		}

		userID, err := cmd(c.Request.Context(), domain.User{
			Email:    createUserReq.Email,
			Password: createUserReq.Password,
			Address:  createUserReq.Address,
		})
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrUserAlreadyExist):
				c.Status(http.StatusConflict)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusCreated, createUserResponse{UserID: userID})
	}
}
