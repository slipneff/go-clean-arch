package main

import (
	"fmt"

	"github.com/4kayDev/logger"
	"github.com/4kayDev/logger/log"
	"github.com/slipneff/go-clean-arch/internal/di"
	desc "github.com/slipneff/go-clean-arch/internal/generated/proto/user"
	"github.com/slipneff/go-clean-arch/internal/utils/config"
	"github.com/slipneff/go-clean-arch/internal/utils/flags"
	"google.golang.org/grpc/reflection"
)

func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	logger.ConfigureZeroLogger()
	container := di.New(cfg)

	desc.RegisterUserServiceServer(container.GetGRPCServer(), container.GetRpcServer())
	log.Info("PermissionService successfully registered")

	grpcServer := container.GetGRPCServer()
	reflection.Register(grpcServer)
	log.Info(fmt.Sprintf("serving GRPC on %s:%d", cfg.Host, cfg.Port))
	err := grpcServer.Serve(*container.GetNetListener())
	if err != nil {
		log.Panic(err, "Error while serving grpcServer")
		panic(err)
	}
}
