package campaing

import (
	"EmailN/internal/contract"
	"EmailN/internalerrors"
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
		return "", internalerrors.ErrInternal
	}

	return campaing.ID, nil
}
