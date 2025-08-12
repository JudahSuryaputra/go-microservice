package user

import (
	"context"
	"go-microservice/internal/repository"
	"go-microservice/internal/shared"
	"go-microservice/internal/shared/dto"
)

type (
	Service interface {
		GetUserByID(ctx context.Context, userID string) (resp *dto.GetUserByIDResponse, err error)
		CreateUser(ctx context.Context, request *dto.CreateUserRequest) (resp *dto.CreateUserResponse, err error)
	}

	implUser struct {
		deps shared.Deps

		repo repository.Holder
	}
)

func NewUserService(deps shared.Deps, repo repository.Holder) Service {
	return &implUser{deps: deps, repo: repo}
}
