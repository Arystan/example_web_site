package db

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	UserCol *mongo.Collection
	ID      *snowflake.Node
}

func NewUserRepository(client *mongo.Client, node *snowflake.Node) *userRepository {
	userCol := client.Database(databasePizza).Collection(collectionUsers)
	return &userRepository{
		UserCol: userCol,
		ID:      node,
	}
}


type User struct {
	ID        int64  `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
}


func (r *userRepository) GetUsers() ([]User, error) {
	cursor, err := r.UserCol.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	users := make([]User, 0)
	for cursor.Next(context.TODO()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}