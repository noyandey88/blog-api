package cmd

import (
	"fmt"
	"os"

	"github.com/noyandey88/blog-api/config"
	"github.com/noyandey88/blog-api/infra/db"
	"github.com/noyandey88/blog-api/internal/user"
	"github.com/noyandey88/blog-api/repo"
	"github.com/noyandey88/blog-api/rest"
	userHandler "github.com/noyandey88/blog-api/rest/handlers/user"
)

func Serve() {
	cfg := config.GetConfig()

	dbCon, err := db.NewConnection(cfg.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	userRepo := repo.NewUserRepo(dbCon)

	// domains
	userService := user.NewService(userRepo)

	//middlewares
	// middlewares := middlewares.NewMiddlewares(cfg)

	// handlers
	userHandler := userHandler.NewHandler(userService, cfg)

	server := rest.NewServer(cfg, userHandler)

	server.Start()
}
