package rpc

import (
	"context"
	"fmt"

	"github.com/4kayDev/logger/log"
	desc "github.com/slipneff/go-clean-arch/internal/generated/proto/user"
	"github.com/slipneff/go-clean-arch/internal/pkg/storage/sql"
	"github.com/slipneff/go-clean-arch/internal/utils/json"
)

func (s *Server) GetUserById(ctx context.Context, req *desc.GetUserByIdRequest) (*desc.UserResponse, error) {
	log.Debug(fmt.Sprintf("GetUserById Request, %s", json.ToColorJson(req)))

	response, err := s.userService.GetUserById(ctx, req.UserId)
	if err != nil {
		if err == sql.ErrEntityNotExist {
			return &desc.UserResponse{Error: desc.UserErrorCode_USER_ERROR_CODE_NOT_FOUND}, nil
		}
		log.Error(err, "Error in service GetUserById", "[RPC] GetUserById")
		return nil, err
	}

	return &desc.UserResponse{
		User: &desc.User{
			UserId: response.Id.String(),
			Name:   response.Name,
		},
	}, nil
}
