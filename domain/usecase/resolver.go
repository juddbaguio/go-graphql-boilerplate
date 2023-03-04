package usecase

//go:generate go run github.com/99designs/gqlgen generate

import "go-graphql-boilerplate/domain/ports"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todoRepo ports.ITodoRepository
	userRepo ports.IUserRepository
}
