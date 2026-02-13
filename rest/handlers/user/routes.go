package user

import (
	"net/http"

	"github.com/noyandey88/blog-api/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /user/sign-up", manager.With(
		http.HandlerFunc(h.CreateUser),
	))

	mux.Handle("POST /user/login", manager.With(
		http.HandlerFunc(h.Login),
	))
}
