package tag

import (
	"net/http"
	"strconv"

	"github.com/noyandey88/blog-api/utils"
)

func (h *Handler) GetTag(w http.ResponseWriter, r *http.Request) {
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

	utils.SendData(w, true, "Tag loaded successfully", tag, http.StatusOK)
}
