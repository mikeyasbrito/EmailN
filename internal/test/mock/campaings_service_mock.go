package mock

import (
	"EmailN/internal/contract"

	"github.com/stretchr/testify/mock"
)

type CampaingServiceMock struct {
	mock.Mock
}

func (r *CampaingServiceMock) Create(newCampaing contract.NewCampaing) (string, error) {
	args := r.Called(newCampaing)
	return args.String(0), args.Error(1)
}

func (r *CampaingServiceMock) GetBy(id string) (*contract.CampaingResponse, error) {
	//args := r.Called(id)
	return nil, nil
}
