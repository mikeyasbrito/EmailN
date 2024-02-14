package endpoints

import (
	"EmailN/internal/contract"
	internalmock "EmailN/internal/test/mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaingsGetById_should_return_campaing(t *testing.T) {
	assert := assert.New(t)
	campaing := contract.CampaingResponse{
		ID:      "343",
		Name:    "Test",
		Content: "Hi!",
		Status:  "Pending",
	}
	service := new(internalmock.CampaingServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaing, nil)
	handler := Handler{CampaingService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaingGetById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaing.ID, response.(*contract.CampaingResponse).ID)
	assert.Equal(campaing.Name, response.(*contract.CampaingResponse).Name)
	assert.Equal(campaing.Content, response.(*contract.CampaingResponse).Content)
}

func Test_CampaingsGetById_should_return_error_when_something_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaingServiceMock)
	errExpected := errors.New("Error")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaingService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, errReturned := handler.CampaingGetById(rr, req)

	assert.Equal(errExpected.Error(), errReturned.Error())

}
