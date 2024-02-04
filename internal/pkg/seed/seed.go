package seed

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		{
			Name: "User-GrandAdmin",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, uuid.MustParse("3cb3253a-0a9f-4fc2-9fc7-b15c5cd45ff0"), "Admin")
			},
		},
	}
}
