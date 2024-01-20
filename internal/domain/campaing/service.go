package campaing

import (
	"EmailN/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaing) (string, error) {

	campaing, _ := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	s.Repository.Save(campaing)
	return campaing.ID, nil
}
