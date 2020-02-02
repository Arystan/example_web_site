package main

import (
	"context"
	"example_web_site/handlers"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoClient := initMongoClient()
	node := initSnowflakeNode(1)
	router := handlers.InitRouter(mongoClient, node)
	if err := router.Run(":8888"); err != nil {
		logrus.Fatal(err)
	}
}

func initMongoClient() *mongo.Client {
	mongoURI := fmt.Sprintf("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(mongoURI)

	logrus.Info("Connecting to Mongo at:", "localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("Connected to Mongo:", "localhost:27017")
	return client
}

func initSnowflakeNode(nodeId int64) *snowflake.Node {
	snowflakeNode, err := snowflake.NewNode(nodeId)
	if err != nil {
		logrus.Fatal(err)
	}
	return snowflakeNode
}