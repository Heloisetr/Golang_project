package main

import (
	"ads/internal/conf"
	"ads/internal/infrastructure/database"
	"ads/internal/transport/http"
	"ads/internal/usecase"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

func initRoute(router *gin.Engine, client *mongo.Client) {
	router.POST("/ad", http.AuthorizeMiddleware(), http.CreateAdHandler(usecase.CreateAd(client)))
	router.GET("/ad/get/:ad_id", http.GetAdHandler(usecase.GetAd(client)))
	router.DELETE("/ad/delete/:ad_id", http.AuthorizeMiddleware(), http.DeleteAdHandler(usecase.DeleteAd(client)))
	router.PATCH("/ad/update/:ad_id", http.AuthorizeMiddleware(), http.UpdateAdHandler(usecase.UpdateAd(client)))
	router.GET("/ad/get_all/:user_id", http.AuthorizeMiddleware(), http.GetAllsAdHandler(usecase.GetAllAd(client)))
	router.GET("/ad/get_by_keys/:keyword", http.GetByKeysAdHandler(usecase.GetByKeysAd(client)))
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
		logrus.Fatal("error while runnig the router")
	}
}
