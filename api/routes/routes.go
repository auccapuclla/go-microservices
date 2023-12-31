package routes

import (
	"microservices/api/middlewares"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, routes []*Route) {
	for _, route := range routes {
		router.HandleFunc(route.Path, middlewares.LogRequests(route.Handler)).
			Methods(route.Method)
	}
}

func WithCORS(router *mux.Router) http.Handler {
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Accept", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete})
	return handlers.CORS(header, origins, methods)(router)
}
