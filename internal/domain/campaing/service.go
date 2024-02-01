package campaing

import (
	"EmailN/internal/contract"
	"EmailN/internalerrors"
)

type Service interface {
	Create(newCampaing contract.NewCampaing) (string, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaing contract.NewCampaing) (string, error) {

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
