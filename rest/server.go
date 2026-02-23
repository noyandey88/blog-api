package rest

import (
	"fmt"
	"net/http"
	"os"

	"github.com/noyandey88/blog-api/config"
	"github.com/noyandey88/blog-api/rest/handlers/category"
	"github.com/noyandey88/blog-api/rest/handlers/tag"
	"github.com/noyandey88/blog-api/rest/handlers/user"
	"github.com/noyandey88/blog-api/rest/middlewares"
)

type Server struct {
	cfg             *config.Config
	userHandler     *user.Handler
	categoryHandler *category.Handler
	tagHandler      *tag.Handler
}

func NewServer(
	cfg *config.Config,
	userHandler *user.Handler,
	categoryHandler *category.Handler,
	tagHandler *tag.Handler,
) *Server {
	return &Server{
		cfg:             cfg,
		userHandler:     userHandler,
		categoryHandler: categoryHandler,
		tagHandler:      tagHandler,
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

	server.userHandler.RegisterRoutes(mux, manager)
	server.categoryHandler.RegisterRoutes(mux, manager)

	addr := fmt.Sprintf(":%d", server.cfg.HttpPort)

	fmt.Println("Server is running on port", addr)
	err := http.ListenAndServe(addr, wrappedMux)

	if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}

}
