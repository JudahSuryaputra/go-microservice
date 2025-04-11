package repository

import (
	"go.uber.org/dig"
)

type Holder struct {
	dig.In

	CacheRepository CacheRepository
	DataRepository  DataRepository
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewDataRepository); err != nil {
		return err
	}
	if err := container.Provide(NewCacheRepository); err != nil {
		return err
	}
	return nil
}
