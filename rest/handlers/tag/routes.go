package tag

import (
	"net/http"

	"github.com/noyandey88/blog-api/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /tags/all", manager.With(
		http.HandlerFunc(h.GetTags),
	))

	mux.Handle("POST /tags/create", manager.With(
		http.HandlerFunc(h.CreateTag),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("GET /tags/get/{id}", manager.With(
		http.HandlerFunc(h.GetTag),
	))

	mux.Handle("PUT /tags/update", manager.With(
		http.HandlerFunc(h.UpdateTag),
		h.middlewares.AuthenticateJWT,
	))

	mux.Handle("DELETE /tags/delete/{id}", manager.With(
		http.HandlerFunc(h.DeleteTag),
		h.middlewares.AuthenticateJWT,
	))
}
