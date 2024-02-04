package rpc

import (
	"context"
	"fmt"

	"github.com/4kayDev/logger/log"
	desc "github.com/slipneff/go-clean-arch/internal/generated/proto/user"
	"github.com/slipneff/go-clean-arch/internal/pkg/storage/sql"
	"github.com/slipneff/go-clean-arch/internal/utils/json"
)

func (s *Server) DeleteUser(ctx context.Context, req *desc.DeleteUserRequest) (*desc.UserResponse, error) {
	log.Debug(fmt.Sprintf("DeleteUser Request, %s", json.ToColorJson(req)))

	response, err := s.userService.DeleteUser(ctx, req.UserId)
	if err != nil {
		if err == sql.ErrEntityNotExist {
			return &desc.UserResponse{Error: desc.UserErrorCode_USER_ERROR_CODE_NOT_FOUND}, nil
		}
		log.Error(err, "Error in service DeleteUser", "[RPC] DeleteUser")
		return nil, err
	}

	return &desc.UserResponse{
		User: &desc.User{
			UserId: response.Id.String(),
			Name:   response.Name,
		},
	}, nil
}
