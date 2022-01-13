package main

import (
	"account/internal/conf"
	"account/internal/infrastructure/db"
	"account/internal/infrastructure/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error Loading Env: %v", err)
	}

	config, err := conf.InitConfig(fmt.Sprintf("conf/%s.yaml", os.Getenv("ENV")))
	if err != nil {
		logrus.Fatalf("Error initializing Conf: %v", err)
	}

	client := db.ConnectDatabase()

	router := gin.Default()

	routes.InitRoute(router, client)

	err = router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
	if err != nil {
		logrus.Fatal("error while running the router")
	}
}
