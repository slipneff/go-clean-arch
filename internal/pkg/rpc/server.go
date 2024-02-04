package rpc

import (
	desc "github.com/slipneff/go-clean-arch/internal/generated/proto/user"
	"github.com/slipneff/go-clean-arch/internal/pkg/service/user"
)

type Server struct {
	desc.UnimplementedUserServiceServer

	userService *user.UserService
}

func New(userService *user.UserService) *Server {
	return &Server{
		userService: userService,
	}
}
