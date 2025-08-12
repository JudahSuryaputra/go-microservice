package user

import (
	"context"
	"go-microservice/internal/shared/dto"
)

func (s *implUser) CreateUser(ctx context.Context, request *dto.CreateUserRequest) (resp *dto.CreateUserResponse, err error) {
	return resp, nil
}
