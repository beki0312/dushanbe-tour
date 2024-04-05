package repository

import (
	"dushanbe_tour/internal/models"
	"dushanbe_tour/internal/pkg/db"
	"go.uber.org/fx"
)

var NewRepository = fx.Provide(newRepository)

type IRepository interface {
	GeLanguage() (lang []models.Language, errs error)
	GeInfrastructure(IdLanguage int64) (infrastruction []models.Infrastructure, errs error)
	GeCategoryInfrastructure() (infrastructure []models.InfrastructureAll, errs error)
}

type dependencies struct {
	fx.In
	Postgres db.IPostgres
}

type repository struct {
	Postgres db.IPostgres
}

func newRepository(dp dependencies) IRepository {
	return &repository{
		Postgres: dp.Postgres,
	}
}
