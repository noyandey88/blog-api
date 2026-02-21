package category

import "github.com/noyandey88/blog-api/config"

type Handler struct {
	cfg *config.Config
	svc Service
}

func NewHandler(svc Service, cfg *config.Config) *Handler {
	return &Handler{
		cfg: cfg,
		svc: svc,
	}
}
