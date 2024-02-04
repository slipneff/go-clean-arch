package testsql

import (
	"testing"

	"github.com/slipneff/go-clean-arch/internal/pkg/model"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func MustCreateUser(t *testing.T, db *gorm.DB, m model.User) *model.User {
	err := db.Create(&m).Error
	require.NoError(t, err)
	return &m
}

func MustFindUser(t *testing.T, db *gorm.DB, id string) *model.User {
	user := model.User{}

	err := db.Model(&user).First(&user, "id = ?", id).Error
	require.NoError(t, err)

	return &user
}

func MustGetErrNotFoundUser(t *testing.T, db *gorm.DB, id string) {
	user := model.User{}

	err := db.Model(&user).First(&user, "id = ?", id).Error

	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
}
