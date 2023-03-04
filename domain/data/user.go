package data

import (
	"context"
	"go-graphql-boilerplate/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	client *mongo.Client
}

func NewUserRepo(client *mongo.Client) *userRepo {
	return &userRepo{
		client: client,
	}
}

func (u *userRepo) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	panic("not implemented yet")
}
