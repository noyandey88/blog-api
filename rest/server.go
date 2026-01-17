package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/noyandey88/blog-api/config"
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
	mux := http.NewServeMux()

	addr := fmt.Sprintf(":%d", server.cfg.HttpPort)

	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, mux)

	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}

}
