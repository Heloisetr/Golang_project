package routes

import (
	"account/internal/transport/http"
	"account/internal/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoute(router *gin.Engine, client *mongo.Client) {
	router.POST("/account", http.CreateAccountHandler(usecase.CreateAccount(client)))
	router.POST("/login", http.LoginHandler(usecase.Login(client)))
	router.GET("/account/:account_id", http.AuthorizeMiddleware(), http.GetAccountHandler(usecase.GetAccount(client)))
	router.DELETE("/account/:account_id", http.AuthorizeMiddleware(), http.DeleteAccountHandler(usecase.DeleteAccount(client)))
	router.POST("/balance", http.AuthorizeMiddleware(), http.AddBalanceHandler(usecase.AddBalance(client)))
	router.PATCH("/update", http.AuthorizeMiddleware(), http.UpdateAccountHandler(usecase.UpdateAccount(client)))
	router.GET("/ping", func(c *gin.Context) { c.JSON(200, "pong") })
}
