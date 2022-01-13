package main

import (
	"epitech/deliveats/user/internal/conf"
	"epitech/deliveats/user/internal/infrastructure/order"
	"epitech/deliveats/user/internal/infrastructure/user"
	"epitech/deliveats/user/internal/transport/http"
	"epitech/deliveats/user/internal/usecase"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func initRoute(router *gin.Engine, userStore user.Store, orderFetcher order.Fetcher) {
	router.POST("/user", http.CreateUserHandler(usecase.CreateUser(userStore)))
	router.GET("/user", http.AuthorizeMiddleware(), http.GetUserHandler(usecase.GetUser(userStore)))
	router.GET("/user/address", http.AuthorizeMiddleware(), http.GetUserAddressHandler(usecase.GetUserAddress(userStore)))
	router.DELETE("/user", http.AuthorizeMiddleware(), http.DeleteUserHandler(usecase.DeleteUser(userStore, orderFetcher)))
	router.GET("/ping", func(c *gin.Context) { c.JSON(200, "pong") })
}

func initConfig(configFilePath string) (conf.Configuration, error) {
	f, err := os.Open(configFilePath)
	if err != nil {
		return conf.Configuration{}, errors.Wrap(err, "unable to open configuration file")
	}
	defer func() { _ = f.Close() }()

	var config = &conf.Configuration{}
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		return conf.Configuration{}, errors.Wrap(err, "unable to load configuration")
	}
	return *config, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env file: %v", err)
	}

	config, err := initConfig(fmt.Sprintf("conf/%s.yaml", os.Getenv("ENV")))
	if err != nil {
		logrus.Fatalf("error initializing configuration: %v", err)
	}

	userStore := user.NewInMemory()
	orderFetcher := order.NewAPI(config.OrderService)

	router := gin.Default()

	initRoute(router, userStore, orderFetcher)

	err = router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		logrus.Fatal("error while running the router")
	}
}
