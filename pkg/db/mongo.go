package db

import (
	"context"
	"log"
	"time"

	"github.com/R3iwan/soundScroll/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(cfg config.MongoConfig) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("MongoDB Connected!")
	return client, nil
}
