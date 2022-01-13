package http

import (
	"epitech/deliveats/user/domain"
	"epitech/deliveats/user/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(cmd usecase.GetUserCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := cmd(c.Request.Context(), c.GetString("user_id"))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrUserNotFound):
				c.Status(http.StatusNotFound)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
