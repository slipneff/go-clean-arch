package sql

import (
	"context"
	"errors"

	"github.com/slipneff/go-clean-arch/internal/pkg/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (s *Storage) CreateUser(ctx context.Context, m model.User) (*model.User, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Create(&m).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrEntityExist
		}
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return nil, ErrForeignKey
		}
		return nil, err
	}

	return &m, nil
}

func (s *Storage) GetUserById(ctx context.Context, id string) (*model.User, error) {
	var currentUser model.User
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	result := tr.First(&currentUser, "id = ?", id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return nil, ErrEntityNotExist
		}
		return nil, result.Error
	}

	return &currentUser, nil
}
func (s *Storage) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrEntityNotExist
		}
		return nil, err
	}
	return users, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	result := tr.Model(&user).Updates(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			return nil, ErrForeignKey
		}
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrEntityNotExist
	}

	return &user, nil
}

func (s *Storage) DeleteUser(ctx context.Context, id string) (*model.User, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var currentUser model.User
	result := tr.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&currentUser)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, ErrEntityNotExist
	}

	return &currentUser, nil
}
