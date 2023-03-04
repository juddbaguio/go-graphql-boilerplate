package ports

import (
	"context"
	"go-graphql-boilerplate/domain/model"
)

type ITodoRepository interface {
	CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	Todos(ctx context.Context) ([]*model.Todo, error)
}
