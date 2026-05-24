package users_service

import (
	"context"
	"github.com/yunishuseynov99/go_todolist/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	createUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(usersRepository UsersRepository) *UsersService {
	return &UsersService{usersRepository: usersRepository}
}
