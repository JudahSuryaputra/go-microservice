package repository

import (
	"go-microservice/internal/entity"
	"go-microservice/internal/shared"
)

type (
	DataRepository interface {
		GetDataById(id string) (data *entity.Data, err error)
	}

	implData struct {
		deps shared.Deps
	}
)

func NewDataRepository(deps shared.Deps) DataRepository {
	return &implData{deps: deps}
}

func (r *implData) GetDataById(id string) (data *entity.Data, err error) {
	err = r.deps.DB.Where("id = ? AND deleted_at IS NULL", id).First(data).Error
	if err != nil {
		return nil, err
	}
	return
}
