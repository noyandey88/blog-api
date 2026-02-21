package category

import (
	"net/http"

	"github.com/noyandey88/blog-api/utils"
)

func (h *Handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categoryList, err := h.svc.List()

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Data loaded successfully", categoryList, http.StatusOK)

}
