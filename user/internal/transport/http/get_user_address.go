package http

import (
	"epitech/deliveats/user/domain"
	"epitech/deliveats/user/internal/usecase"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserAddressHandler(cmd usecase.GetUserAddressCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")

		userAddress, err := cmd(c.Request.Context(), userID)
		fmt.Println(userAddress)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrUserNotFound):
				c.Status(http.StatusNotFound)
			default:
				c.Status(http.StatusInternalServerError)
			}
			return
		}
		c.JSON(http.StatusOK, userAddress)
	}
}
