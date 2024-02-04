package rpc

import (
	"context"
	"fmt"

	"github.com/4kayDev/logger/log"
	"github.com/google/uuid"
	desc "github.com/slipneff/go-clean-arch/internal/generated/proto/user"
	"github.com/slipneff/go-clean-arch/internal/pkg/model"
	"github.com/slipneff/go-clean-arch/internal/pkg/storage/sql"
	"github.com/slipneff/go-clean-arch/internal/utils/json"
)

func (s *Server) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*desc.UserResponse, error) {
	log.Debug(fmt.Sprintf("CreateUser Request, %s", json.ToColorJson(req)))
	if req == nil {
		return &desc.UserResponse{Error: desc.UserErrorCode_USER_ERROR_CODE_VALIDATION}, nil
	}
	var userId uuid.UUID
	if req.User.UserId == "" {
		userId = uuid.Nil
	} else if validateUUID(req.User.UserId, &userId) != nil {
		return &desc.UserResponse{Error: desc.UserErrorCode_USER_ERROR_CODE_VALIDATION}, nil
	}

	response, err := s.userService.CreateUser(ctx, model.User{Id: userId, Name: req.User.Name})
	if err != nil {
		if err == sql.ErrEntityExist {
			return &desc.UserResponse{Error: desc.UserErrorCode_USER_ERROR_CODE_ALREADY_EXIST}, nil
		}
		log.Error(err, "Error in service CreateUser", "[RPC] CreateUser")
		return nil, err
	}
	return &desc.UserResponse{
		User: &desc.User{
			UserId: response.Id.String(),
			Name:   response.Name,
		},
	}, nil
}
