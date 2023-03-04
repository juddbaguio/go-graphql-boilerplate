package infra

import (
	"context"
	"errors"
	"go-graphql-boilerplate/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(ctx context.Context, cfg *config.Mongo) (*mongo.Client, error) {
	if cfg == nil {
		return nil, errors.New("mongo config is missing")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))

	if err != nil {
		return nil, err
	}

	return client, nil
}
