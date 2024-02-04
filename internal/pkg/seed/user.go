package seed

import (
	"github.com/google/uuid"
	"github.com/slipneff/go-clean-arch/internal/pkg/model"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, id uuid.UUID, name string) error {
	return db.Create(&model.User{
		Id:   id,
		Name: name,
	}).Error
}
