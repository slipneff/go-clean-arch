package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/slipneff/go-clean-arch/internal/utils/json"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Name      string     `gorm:"unique"`
	CreatedAt *time.Time `gorm:"column:CREATED_AT;autoCreateTime"`
	DeletedAt *time.Time `gorm:"column:DELETED_AT;autoDeleteTime"`
}

func (u *User) String() string {
	return json.ToColorJson(u)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Id == uuid.Nil {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		u.Id = id
	}
	return nil
}
