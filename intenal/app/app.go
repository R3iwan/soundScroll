package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/R3iwan/soundScroll/pkg/config"
	"github.com/R3iwan/soundScroll/pkg/db"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config %v", err)
	}

	postgresDB, err := db.ConnectPostgres(cfg.PostgresConfig)
	if err != nil {
		log.Fatalf("Error connecting Postgres %v", err)
	}
	defer postgresDB.Close()

	mongoClient, err := db.ConnectMongo(cfg.MongoConfig)
	if err != nil {
		log.Fatalf("Error connecting Mongo %v", err)
	}
	defer mongoClient.Disconnect(context.TODO())

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)

	router.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "route works fine",
		})
	})

	if err := router.Run(fmt.Sprintf(":%s", cfg.ServerConfig.Port)); err != nil {
		log.Fatalf("Error occurred during running server: %v", err)
	}
}
