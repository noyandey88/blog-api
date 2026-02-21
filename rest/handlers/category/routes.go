package category

import (
	"net/http"

	"github.com/noyandey88/blog-api/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /categories/get", manager.With(
		http.HandlerFunc(h.GetCategories),
	))

	mux.Handle("POST /categories/create", manager.With(
		http.HandlerFunc(h.CreateCategory),
	))

	mux.Handle("GET /categories/get/{id}", manager.With(
		http.HandlerFunc(h.GetCategory),
	))

	mux.Handle("PUT /categories/update/{id}", manager.With(
		http.HandlerFunc(h.UpdateCategory),
	))

	mux.Handle("DELETE /categories/delete/{id}", manager.With(
		http.HandlerFunc(),
	))
}
