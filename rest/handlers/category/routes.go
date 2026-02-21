package category

import (
	"net/http"

	"github.com/noyandey88/blog-api/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /categories/get", manager.With(
		http.HandlerFunc(h.GetCategories),
	))
}
