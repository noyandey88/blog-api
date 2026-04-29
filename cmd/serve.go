package cmd

import (
	"fmt"
	"os"

	"github.com/noyandey88/blog-api/config"
	"github.com/noyandey88/blog-api/infra/db"
	"github.com/noyandey88/blog-api/internal/category"
	"github.com/noyandey88/blog-api/internal/tag"
	"github.com/noyandey88/blog-api/internal/user"
	"github.com/noyandey88/blog-api/repo"
	"github.com/noyandey88/blog-api/rest"
	categoryHandler "github.com/noyandey88/blog-api/rest/handlers/category"
	tagHandler "github.com/noyandey88/blog-api/rest/handlers/tag"
	userHandler "github.com/noyandey88/blog-api/rest/handlers/user"
	"github.com/noyandey88/blog-api/rest/middlewares"
)

func Serve() {
	cfg := config.GetConfig()

	dbCon, err := db.NewConnection(cfg.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	userRepo := repo.NewUserRepo(dbCon)
	categoryRepo := repo.NewCategoryRepo(dbCon)
	tagRepo := repo.NewTagRepo(dbCon)

	// domains
	userService := user.NewService(userRepo)
	categoryService := category.NewService(categoryRepo)
	tagService := tag.NewService(tagRepo)

	//middlewares
	middlewares := middlewares.NewMiddlewares(cfg)

	// handlers
	userHandler := userHandler.NewHandler(userService, cfg)
	categoryHandler := categoryHandler.NewHandler(middlewares, categoryService)
	tagHandler := tagHandler.NewHandler(middlewares, tagService)

	server := rest.NewServer(
		cfg,
		userHandler,
		categoryHandler,
		tagHandler,
	)

	server.Start()
}
