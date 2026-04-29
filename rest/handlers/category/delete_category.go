package category

import (
	"net/http"
	"strconv"

	"github.com/noyandey88/blog-api/utils"
)

func (h *Handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.PathValue("id")

	cId, err := strconv.Atoi(categoryId)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	category, err := h.svc.Get(cId)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	if category == nil {
		utils.SendError(w, false, "Category not found", nil, http.StatusNotFound)
		return
	}

	err = h.svc.Delete(cId)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Category deleted successfully", nil, http.StatusOK)
}
