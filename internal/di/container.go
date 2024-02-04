package di

import (
	"fmt"
	"net"

	"github.com/slipneff/go-clean-arch/internal/pkg/rpc"
	"github.com/slipneff/go-clean-arch/internal/pkg/service/user"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/slipneff/go-clean-arch/internal/pkg/storage/sql"
	"github.com/slipneff/go-clean-arch/internal/utils/config"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Container struct {
	cfg *config.Config

	netListener *net.Listener
	grpcServer  *grpc.Server

	db                 *gorm.DB
	transactionManager trm.Manager

	storage      *sql.Storage
	userService *user.UserService
	server       *rpc.Server
}

func New(cfg *config.Config) *Container {
	return &Container{cfg: cfg}
}

func (c *Container) GetNetListener() *net.Listener {
	return get(&c.netListener, func() *net.Listener {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", c.cfg.Port))
		if err != nil {
			panic(err)
		}

		return &listener
	})
}

func (c *Container) GetGRPCServer() *grpc.Server {
	return get(&c.grpcServer, func() *grpc.Server {
		grpcServer := grpc.NewServer()
		return grpcServer
	})
}
func (c *Container) GetPostgresDB() *sql.Storage {
	return get(&c.storage, func() *sql.Storage {
		return sql.New(c.GetDB(), trmgorm.DefaultCtxGetter)
	})
}

func (c *Container) GetDB() *gorm.DB {
	return get(&c.db, func() *gorm.DB {
		return sql.MustNewSQLite(c.cfg)
	})
}

func (c *Container) GetTransactionManager() trm.Manager {
	return get(&c.transactionManager, func() trm.Manager {
		return manager.Must(trmgorm.NewDefaultFactory(c.db))
	})
}

func (c *Container) GetUserService() *user.UserService {
	return get(&c.userService, func() *user.UserService {
		return user.NewUserService(c.GetPostgresDB())
	})
}

func (c *Container) GetRpcServer() *rpc.Server {
	return get(&c.server, func() *rpc.Server {
		return rpc.New(c.GetUserService())
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}

	*obj = builder()
	return *obj
}
