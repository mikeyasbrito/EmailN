package campaing

import (
	"EmailN/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaing *Campaing) error {
	args := r.Called(campaing)
	return args.Error(0)
}

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaing := contract.NewCampaing{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test1@test.com"},
	}

	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_SaveCampaing(t *testing.T) {
	newCampaing := contract.NewCampaing{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test1@test.com"},
	}

	repositoryMock := new(repositoryMock)
	service := Service{Repository: repositoryMock}
	repositoryMock.On("Save", mock.MatchedBy(func(campaing *Campaing) bool {
		if campaing.Name != newCampaing.Name {
			return false
		} else if campaing.Content != newCampaing.Content {
			return false
		} else if len(campaing.Contacts) != len(newCampaing.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service.Create(newCampaing)

	repositoryMock.AssertExpectations(t)
}
