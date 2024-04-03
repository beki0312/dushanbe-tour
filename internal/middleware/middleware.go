package middleware

import (
	"dushanbe_tour/internal/pkg/service"
	"go.uber.org/fx"
	"net/http"
)

var Module = fx.Provide(NewMiddleware)

type Middleware interface {
	Cors(next http.Handler) http.Handler
}

type Params struct {
	fx.In
	SVC service.IService
}

type middleware struct {
	svc service.IService
}

func NewMiddleware(params Params) Middleware {
	return &middleware{
		svc: params.SVC,
	}
}

func (s *middleware) Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS, PUT, HEAD, TRACE, CONNECT")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")
		if r.Method == "OPTIONS" {
			w.Write([]byte("OPTIONS"))
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
