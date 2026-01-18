package middlewares

import "github.com/noyandey88/blog-api/config"

type Middlewares struct {
	cfg *config.Config
}

func NewMiddlewares(cfg *config.Config) *Middlewares {
	return &Middlewares{
		cfg: cfg,
	}
}
