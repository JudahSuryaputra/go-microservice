package repository

import (
	"go-microservice/internal/entity"
	"go-microservice/internal/shared"
)

type (
	UserRepository interface {
		GetUserByID(userID string) (user *entity.User, err error)
	}

	implUser struct {
		deps shared.Deps
	}
)

func NewUserRepository(deps shared.Deps) UserRepository {
	return &implUser{deps: deps}
}

func (r *implUser) GetUserByID(userID string) (user *entity.User, err error) {
	err = r.deps.DB.Where("user_id = ? AND deleted_at IS NULL", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	return
}
