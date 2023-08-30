package server

import "net/http"

type ServerConfig struct {
	Addr   string
	Router RouterConfig
}

func NewServer(cfg ServerConfig) http.Server {
	return http.Server{
		Addr:    cfg.Addr,
		Handler: NewRouter(cfg.Router),
	}
}
