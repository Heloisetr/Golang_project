package main

import (
	"fmt"
	"os"
	"transaction/internal/conf"
	"transaction/internal/infrastructure/database"
	"transaction/internal/transport/http"
	"transaction/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

func initRoute(router *gin.Engine, client *mongo.Client) {
	router.POST("/bid", http.AuthorizeMiddleware(), http.CreateBidHandler(usecase.CreateBid(client)))
	router.GET("/bid/:bid_id", http.AuthorizeMiddleware(), http.GetBidHandler(usecase.GetBid(client)))
	router.PATCH("/bid/update/:bid_id", http.AuthorizeMiddleware(), http.UpdateBidHandler(usecase.UpdateBid(client)))
	router.DELETE("/bid/delete/:bid_id", http.AuthorizeMiddleware(), http.DeleteBidHandler(usecase.DeleteBid(client)))
	router.GET("/bids", http.AuthorizeMiddleware(), http.GetAllBidsHandler(usecase.GetAllBids(client)))
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

	client := database.ConnectToDatabase()

	router := gin.Default()

	initRoute(router, client)

	err = router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		logrus.Fatal("error while running the router")
	}
}
