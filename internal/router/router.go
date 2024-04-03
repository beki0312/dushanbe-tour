package router

import (
	"context"
	"dushanbe_tour/internal/handlers"
	"dushanbe_tour/internal/middleware"
	"dushanbe_tour/internal/pkg/service"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log"
	"net/http"
)

var EntryPoint = fx.Options(
	fx.Invoke(
		NewRouter,
	),
)

type dependencies struct {
	fx.In
	Lifecycle fx.Lifecycle
	//Config    config.IConfig
	SVC        service.IService
	Handler    handlers.IHandler
	Middleware middleware.Middleware
}

func NewRouter(d dependencies) {
	server := mux.NewRouter()
	mainRoute := server.PathPrefix("/api").Subrouter()
	generalMiddleware := []mux.MiddlewareFunc{
		d.Middleware.Cors,
	}
	mainRoute.Use(generalMiddleware...)
	courseRoute := mainRoute.PathPrefix("/v1").Subrouter()

	log.Println("courseRoute=======> ", courseRoute)

	srv := http.Server{
		Addr:    "127.0.5.24:53488", //os.Getenv("APP_PORT"), //d.SVC.ConfigInstance().GetString("api.server.port"),
		Handler: server,
	}
	log.Println("srv  ", srv.Addr)
	d.Lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				//d.SVC.LoggerInstance().Info("Application started")
				log.Println("Application started")
				go srv.ListenAndServe()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				//d.SVC.LoggerInstance().Info("Application stopped")
				log.Println("Application stopped")
				return srv.Shutdown(ctx)
			},
		},
	)
}
