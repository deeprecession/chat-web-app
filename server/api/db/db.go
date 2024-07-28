package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoStorage struct {
	Client *mongo.Client
}

func NewMongoStorage(uri string) (MongoStorage, error) {
	client, err := createMongoClient(uri)
	if err != nil {
		return MongoStorage{}, err
	}

	storage := MongoStorage{Client: client}

	err = storage.initScheme()
	if err != nil {
		return MongoStorage{}, err
	}

	return storage, nil
}

func createMongoClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (storage *MongoStorage) initScheme() error {
	err := storage.createUniqueUsernameIndex()

	return err
}

func (storage *MongoStorage) createUniqueUsernameIndex() error {
	usersCollection := storage.Client.Database("test_db").Collection("users")

	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"name", 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := usersCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		return fmt.Errorf("failed to create unique index on username: %w", err)
	}

	return nil
}
