package cmd

import (
	"fmt"
	"os"

	"github.com/noyandey88/blog-api/config"
	"github.com/noyandey88/blog-api/infra/db"
	"github.com/noyandey88/blog-api/rest"
)

func Serve() {
	cfg := config.GetConfig()

	_, err := db.NewConnection(cfg.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	server := rest.NewServer(cfg)

	server.Start()
}
