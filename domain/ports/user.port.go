package ports

import (
	"context"
	"go-graphql-boilerplate/domain/model"
)

type IUserRepository interface {
	FindUserByID(ctx context.Context, id string) (*model.User, error)
}
