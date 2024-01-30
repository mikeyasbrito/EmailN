package endpoints

import (
	"net/http"
)

func (h *Handler) CampaingGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaings, err := h.CampaingService.Repository.Get()
	return campaings, 200, err

}
