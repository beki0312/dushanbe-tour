package handlers

import (
	"dushanbe_tour/internal/pkg/service"
	"go.uber.org/fx"
)

var NewHandler = fx.Provide(newHandler)

type IHandler interface {
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
