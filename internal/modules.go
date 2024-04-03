package internal

import (
	"dushanbe_tour/internal/handlers"
	"dushanbe_tour/internal/middleware"
	"dushanbe_tour/internal/pkg/db"
	"dushanbe_tour/internal/pkg/repository"
	"dushanbe_tour/internal/pkg/service"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	service.NewService,
	handlers.NewHandler,
	repository.NewRepository,
	db.NewPostgres,
	middleware.Module,
)
