package tag

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/utils"
)

type UpdateTagReq struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) UpdateTag(w http.ResponseWriter, r *http.Request) {
	var req UpdateTagReq

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	tag, err := h.svc.Get(req.ID)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	if tag == nil {
		utils.SendError(w, false, "Tag not found", nil, http.StatusNotFound)
		return
	}

	updatedTag, err := h.svc.Update(domain.Tag{
		ID:        req.ID,
		Name:      req.Name,
		UpdatedAt: time.Now().Unix(),
	})

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Category updated successfully", updatedTag, http.StatusOK)

}
