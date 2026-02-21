package category

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/utils"
)

type UpdateCategoryReq struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

func (h *Handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
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

	var req UpdateCategoryReq

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	req.ID = cId

	updatedCategory, err := h.svc.Update(domain.Category{
		ID:          cId,
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
	})

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Category updated successfully", updatedCategory, http.StatusOK)

}
