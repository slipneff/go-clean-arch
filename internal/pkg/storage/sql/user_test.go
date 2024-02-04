package sql_test

import (
	"context"
	"testing"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/slipneff/go-clean-arch/internal/pkg/storage/sql"
	"github.com/slipneff/go-clean-arch/internal/test/fake"
	"github.com/slipneff/go-clean-arch/internal/test/testsql"
	"github.com/stretchr/testify/require"
)

func TestStorage_CreateUser(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = sql.MustNewTestDB(t)
		storage = sql.New(db, trmgorm.DefaultCtxGetter)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		fakeUser := fake.User()

		user, err := storage.CreateUser(ctx, fakeUser)
		require.NoError(t, err)

		actual := testsql.MustFindUser(t, db, user.Id.String())
		user.CreatedAt = actual.CreatedAt
		require.Equal(t, user, actual)

	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		user := fake.User()

		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := storage.CreateUser(ctx, user)
		require.Error(t, err)
	})

	t.Run("entity already exists", func(t *testing.T) {
		t.Parallel()
		user := fake.User()

		_, err := storage.CreateUser(ctx, user)
		require.NoError(t, err)

		_, err = storage.CreateUser(ctx, user)
		require.ErrorIs(t, err, sql.ErrEntityExist)
	})

}

func TestStorage_ReadUser(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = sql.MustNewTestDB(t)
		storage = sql.New(db, trmgorm.DefaultCtxGetter)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		fakeUser := fake.User()
		user := testsql.MustCreateUser(t, db, fakeUser)

		actual, err := storage.GetUserById(ctx, user.Id.String())
		require.NoError(t, err)
		user.CreatedAt = actual.CreatedAt
		require.Equal(t, user, actual)

	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		user := fake.User()
		testsql.MustCreateUser(t, db, user)
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := storage.GetUserById(ctx, user.Id.String())
		require.Error(t, err)
	})

	t.Run("entity not found", func(t *testing.T) {
		t.Parallel()
		user := fake.User()

		_, err := storage.GetUserById(ctx, user.Id.String())
		require.ErrorIs(t, err, sql.ErrEntityNotExist)
	})
}

func TestStorage_UpdateUser(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = sql.MustNewTestDB(t)
		storage = sql.New(db, trmgorm.DefaultCtxGetter)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		user := fake.User()

		testsql.MustCreateUser(t, db, user)

		updateUser := fake.User()
		updateUser.Id = user.Id

		updated, err := storage.UpdateUser(ctx, updateUser)
		require.NoError(t, err)

		actual := testsql.MustFindUser(t, db, updated.Id.String())
		updateUser.CreatedAt = actual.CreatedAt
		require.Equal(t, updateUser, *actual)

	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		user := fake.User()
		testsql.MustCreateUser(t, db, user)
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := storage.UpdateUser(ctx, user)
		require.Error(t, err)
	})

	t.Run("error not found", func(t *testing.T) {
		t.Parallel()
		user := fake.User()

		_, err := storage.UpdateUser(ctx, user)
		require.Error(t, err)
	})
}
func TestStorage_DeleteUser(t *testing.T) {
	var (
		ctx     = context.Background()
		db      = sql.MustNewTestDB(t)
		storage = sql.New(db, trmgorm.DefaultCtxGetter)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		user := fake.User()

		testsql.MustCreateUser(t, db, user)

		deleted, err := storage.DeleteUser(ctx, user.Id.String())
		require.NoError(t, err)

		testsql.MustGetErrNotFoundUser(t, db, deleted.Id.String())
	})

	t.Run("error", func(t *testing.T) {
		t.Parallel()
		user := fake.User()

		testsql.MustCreateUser(t, db, user)
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		_, err := storage.DeleteUser(ctx, user.Id.String())
		require.Error(t, err)
	})
	t.Run("entity not found", func(t *testing.T) {
		t.Parallel()
		user := fake.User()

		_, err := storage.DeleteUser(ctx, user.Id.String())

		require.ErrorIs(t, err, sql.ErrEntityNotExist)
	})
}
