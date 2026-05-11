package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:TvJtgWL0ofSUpBCV@cluster0.vp2nhjw.mongodb.net/?appName=Cluster0")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	DB = client.Database("liga")
	return nil
}
