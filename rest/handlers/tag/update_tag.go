package tag

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/utils"
)

type UpdateTagReq struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) UpdateTag(w http.ResponseWriter, r *http.Request) {
	tagId := r.PathValue("id")

	cId, err := strconv.Atoi(tagId)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	tag, err := h.svc.Get(cId)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	if tag == nil {
		utils.SendError(w, false, "Tag not found", nil, http.StatusNotFound)
		return
	}

	var req UpdateTagReq

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	req.ID = cId

	updatedTag, err := h.svc.Update(domain.Tag{
		ID:   cId,
		Name: req.Name,
	})

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Category updated successfully", updatedTag, http.StatusOK)

}
