package server

import (
	"net/http"

	"github.com/RTK-Tickets/alerts/graphql"
	"github.com/go-chi/chi"
)

type RouterConfig struct {
	GraphQLEndpointPath string
}

func NewRouter(cfg RouterConfig) chi.Router {
	r := chi.NewRouter()

	r.Method(http.MethodPost, cfg.GraphQLEndpointPath, graphql.AlertQueryHandler)

	return r
}
