package db

import (
	"github.com/bwmarrin/snowflake"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databasePizza = "pizza"
)

const (
	collectionUsers    = "users"
	collectionProducts = "products"
)

var UserRepo *userRepository

func Init(client *mongo.Client, node *snowflake.Node) {
	UserRepo = NewUserRepository(client, node)
}
