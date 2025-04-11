package data

import "go-microservice/internal/shared"

type (
	Service interface {
	}

	implData struct {
		deps shared.Deps
	}
)

func NewDataService(deps shared.Deps) (s Service, err error) {
	return &implData{deps: deps}, nil
}
