package rpc

import (
	"context"

	"github.com/4kayDev/logger/log"
	desc "github.com/slipneff/go-clean-arch/internal/generated/proto/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAllUser(ctx context.Context, req *desc.GetAllUserRequest) (*desc.GetAllUserResponse, error) {
	users, err := s.userService.GetAllUsers(ctx)
	if err != nil {
		log.Error(err, "Error in service GetAllUsers", "[RPC] GetAllUsers")
		return nil, status.Error(codes.Internal, "Failed to GetAllUsers")
	}

	response := &desc.GetAllUserResponse{}
	for _, v := range users {
		useId := v.Id.String()
		a := &desc.User{
			UserId: useId,
			Name:   v.Name,
		}

		response.Users = append(response.Users, a)
	}
	return response, nil
}
