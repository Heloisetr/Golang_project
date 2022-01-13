package http

import (
	"epitech/deliveats/user/domain"
	"epitech/deliveats/user/internal/usecase"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(cmd usecase.DeleteUserCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := cmd(c.Request.Context(), c.GetString("user_id"))
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrUserNotFound):
				c.Status(http.StatusNotFound)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		c.Status(http.StatusOK)
	}
}
