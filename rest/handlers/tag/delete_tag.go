package tag

import (
	"net/http"
	"strconv"

	"github.com/noyandey88/blog-api/utils"
)

func (h *Handler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	tagId := r.PathValue("id")

	tId, err := strconv.Atoi(tagId)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	tag, err := h.svc.Get(tId)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	if tag == nil {
		utils.SendError(w, false, "Tag not found", nil, http.StatusNotFound)
		return
	}

	err = h.svc.Delete(tId)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Tag deleted successfully", nil, http.StatusOK)
}
