package user

import (
	"context"

	"github.com/slipneff/go-clean-arch/internal/pkg/model"
)

type userStorage interface {
	CreateUser(ctx context.Context, m model.User) (*model.User, error)
	GetUserById(ctx context.Context, adminID string) (*model.User, error)
	UpdateUser(ctx context.Context, admin model.User) (*model.User, error)
	DeleteUser(ctx context.Context, adminID string) (*model.User, error)
	GetAllUsers(ctx context.Context) ([]model.User, error)
}
type UserService struct {
	storage userStorage
}

func NewUserService(storage userStorage) *UserService {
	return &UserService{storage: storage}
}
