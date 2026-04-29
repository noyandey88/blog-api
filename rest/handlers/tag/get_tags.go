package tag

import (
	"net/http"

	"github.com/noyandey88/blog-api/utils"
)

func (h *Handler) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.svc.List()

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Tags Retrived Successfully", tags, http.StatusOK)
}
