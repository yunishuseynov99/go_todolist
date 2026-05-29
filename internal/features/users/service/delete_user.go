package users_service

import (
	"context"
	"fmt"
)

func (s *UsersService) DeleteUser(ctx context.Context, Id int) error {
	if err := s.usersRepository.DeleteUser(ctx, Id); err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	return nil
}
