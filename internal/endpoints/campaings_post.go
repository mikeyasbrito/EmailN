package endpoints

import (
	"EmailN/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaingPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaing
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaingService.Create(request)
	return map[string]string{"id": id}, 201, err

}
