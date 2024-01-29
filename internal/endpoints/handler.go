package endpoints

import "EmailN/internal/domain/campaing"

type Handler struct {
	CampaingService campaing.Service
}
