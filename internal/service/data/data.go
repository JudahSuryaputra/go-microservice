package data

import (
	"go-microservice/internal/repository"
	"go-microservice/internal/shared"
)

type (
	Service interface {
	}

	implData struct {
		deps shared.Deps

		repo repository.Holder
	}
)

func NewDataService(deps shared.Deps, repo repository.Holder) Service {
	return &implData{deps: deps, repo: repo}
}
