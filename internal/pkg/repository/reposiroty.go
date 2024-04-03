package repository

import (
	"dushanbe_tour/internal/pkg/db"
	"go.uber.org/fx"
)

var NewRepository = fx.Provide(newRepository)

type IRepository interface {
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
