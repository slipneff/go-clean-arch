package user

import (
	"context"
	"fmt"

	"github.com/4kayDev/logger/log"
	"github.com/slipneff/go-clean-arch/internal/pkg/model"
	"github.com/slipneff/go-clean-arch/internal/pkg/service"
	"github.com/slipneff/go-clean-arch/internal/pkg/storage/sql"
	"github.com/slipneff/go-clean-arch/internal/utils/json"
)

func (s *UserService) CreateUser(ctx context.Context, m model.User) (*model.User, error) {
	user, err := s.storage.CreateUser(ctx, m)
	if err != nil {
		if err == sql.ErrForeignKey {
			return nil, service.ErrForeignKey
		}
		return nil, err
	}
	log.Debug(fmt.Sprintf("CreateUser service response, %s", json.ToColorJson(user)))
	return user, nil
}

func (s *UserService) GetUserById(ctx context.Context, userID string) (*model.User, error) {
	currentUser, err := s.storage.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	log.Debug(fmt.Sprintf("GetUserById service response, %s", json.ToColorJson(currentUser)))
	return currentUser, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	users, err := s.storage.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	currentUser, err := s.storage.UpdateUser(ctx, user)
	if err != nil {
		if err == sql.ErrForeignKey {
			return nil, service.ErrForeignKey
		}
		return nil, err
	}
	log.Debug(fmt.Sprintf("UpdateUser service response, %s", json.ToColorJson(currentUser)))
	return currentUser, nil
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) (*model.User, error) {
	currentUser, err := s.storage.DeleteUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	log.Debug(fmt.Sprintf("DeleteUser service response, %s", json.ToColorJson(currentUser)))
	return currentUser, nil
}
