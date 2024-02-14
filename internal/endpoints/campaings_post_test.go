package endpoints

import (
	"EmailN/internal/contract"
	internalmock "EmailN/internal/test/mock"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaingsPost_should_save_new_campaing(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaing{
		Name:    "test",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaingServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaing) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		} else {
			return false
		}
	})).Return("34x", nil)
	handler := Handler{CampaingService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaingPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaingsPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaing{
		Name:    "test",
		Content: "Hi everyone",
		Emails:  []string{"teste@teste.com"},
	}
	service := new(internalmock.CampaingServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaingService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaingPost(rr, req)

	assert.NotNil(err)
}
