package tag

import (
	"encoding/json"
	"net/http"

	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/utils"
)

type CreateTagReq struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (h *Handler) CreateTag(w http.ResponseWriter, r *http.Request) {
	var req CreateTagReq

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&req)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	tag, err := h.svc.Create(domain.Tag{
		Name: req.Name,
		Slug: req.Slug,
	})

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Tag created successfully", tag, http.StatusCreated)
}
