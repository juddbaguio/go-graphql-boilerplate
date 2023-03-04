package usecase

import "go-graphql-boilerplate/domain/ports"

func New(userRepo ports.IUserRepository, todoRepo ports.ITodoRepository) *Resolver {
	return &Resolver{
		todoRepo: todoRepo,
		userRepo: userRepo,
	}
}
