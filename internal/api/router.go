package api

import (
	"ClearOTC/internal/api/handler"
	"net/http"
)

type HTTPRouter struct {
	Handler http.Handler
}

func (H HTTPRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	H.Handler.ServeHTTP(w, r)
}

func NewRouter() *HTTPRouter {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handler.HealthCheck)
	return &HTTPRouter{Handler: mux}
}
