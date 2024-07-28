package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserNotExist     = errors.New("user does not exist")
	UserAlreadyExist = errors.New("user already exist")
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"name"`
	Password string             `bson:"password"`
}

func (storage MongoStorage) GetUser(name string) (User, error) {
	usersCollection := storage.Client.Database("test_db").Collection("users")

	var user User
	filterByName := bson.D{{"name", name}}

	err := usersCollection.FindOne(context.TODO(), filterByName).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return User{}, UserNotExist
		}

		return User{}, err
	}

	return user, nil
}

func (storage MongoStorage) InsertUser(user User) error {
	usersCollection := storage.Client.Database("test_db").Collection("users")

	_, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return UserAlreadyExist
		}

		return err
	}

	return nil
}
