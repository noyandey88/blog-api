package category

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/utils"
)

type UpdateCategoryReq struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var req UpdateCategoryReq
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	category, err := h.svc.Get(req.ID)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	if category == nil {
		utils.SendError(w, false, "Category not found", nil, http.StatusNotFound)
		return
	}

	updatedCategory, err := h.svc.Update(domain.Category{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		UpdatedAt:   time.Now().Unix(),
	})

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Category updated successfully", updatedCategory, http.StatusOK)

}
