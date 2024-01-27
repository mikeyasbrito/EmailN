package database

import (
	"EmailN/internal/domain/campaing"
	"errors"
)

type CampaingRepository struct {
	campaings []campaing.Campaing
}

func (c *CampaingRepository) Save(campaing *campaing.Campaing) error {
	c.campaings = append(c.campaings, *campaing)
	return errors.New("an error")
}
