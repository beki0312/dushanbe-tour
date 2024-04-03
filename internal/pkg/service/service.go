package service

import (
	"dushanbe_tour/internal/pkg/repository"
	"go.uber.org/fx"
)

var NewService = fx.Provide(newService)

type IService interface {
	RepositoryInstance() repository.IRepository
}

type dependencies struct {
	fx.In
	Repository repository.IRepository
}

type service struct {
	Repository repository.IRepository
}

func newService(d dependencies) IService {
	return &service{
		d.Repository,
	}
}

func (s service) RepositoryInstance() repository.IRepository {
	return s.Repository
}
