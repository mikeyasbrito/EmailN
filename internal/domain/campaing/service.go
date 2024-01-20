package campaing

import (
	"EmailN/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaing) (string, error) {

	campaing, err := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	if err != nil {
		return "", err
	}
	err = s.Repository.Save(campaing)
	if err != nil {
		return "", err
	}

	return campaing.ID, nil
}
