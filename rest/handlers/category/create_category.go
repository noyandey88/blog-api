package category

import (
	"encoding/json"
	"net/http"

	"github.com/noyandey88/blog-api/domain"
	"github.com/noyandey88/blog-api/utils"
)

type CreateCategoryReq struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

func (h *Handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryReq

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&req)

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusBadRequest)
		return
	}

	createdCategory, err := h.svc.Create(domain.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
	})

	if err != nil {
		utils.SendError(w, false, err.Error(), nil, http.StatusInternalServerError)
		return
	}

	utils.SendData(w, true, "Category created successfully", createdCategory, http.StatusOK)

}
