package data

import (
	"context"
	"go-graphql-boilerplate/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type todoRepo struct {
	client *mongo.Client
}

func NewTodoRepo(client *mongo.Client) *todoRepo {
	return &todoRepo{
		client: client,
	}
}

func (r *todoRepo) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic("not implemented yet")
}

func (r *todoRepo) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic("not implemented yet")
}
