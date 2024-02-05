package campaing

import (
	"EmailN/internalerrors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending = "Pending"
	Started = "Started"
	Done    = "Done"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaing struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
}

func NewCampaing(name string, content string, emails []string) (*Campaing, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	campaing := &Campaing{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}
	err := internalerrors.ValidateStruct(campaing)
	if err == nil {
		return campaing, nil
	}
	return nil, err
}
