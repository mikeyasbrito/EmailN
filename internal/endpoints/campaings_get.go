package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaingGet(w http.ResponseWriter, r *http.Request) {

	render.Status(r, 200)
	render.JSON(w, r, h.CampaingService.Repository.Get())
}
