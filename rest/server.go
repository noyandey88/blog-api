package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/noyandey88/blog-api/config"
	"github.com/noyandey88/blog-api/rest/middlewares"
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (server *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(
		middlewares.Preflight,
		middlewares.Cors,
		middlewares.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	addr := fmt.Sprintf(":%d", server.cfg.HttpPort)

	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}

}
