package handlers

import (
	"dushanbe_tour/internal/pkg/service"
	"go.uber.org/fx"
	"net/http"
)

var NewHandler = fx.Provide(newHandler)

type IHandler interface {
	GetLanguage() http.HandlerFunc
	GetInfrastructure() http.HandlerFunc
	GetCategoryInfrastructure() http.HandlerFunc
}
type dependencies struct {
	fx.In
	SVC service.IService
}
type Handler struct {
	svc service.IService
}

func newHandler(d dependencies) IHandler {
	return Handler{
		svc: d.SVC,
	}
}
