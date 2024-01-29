package database

import (
	"EmailN/internal/domain/campaing"
)

type CampaingRepository struct {
	campaings []campaing.Campaing
}

func (c *CampaingRepository) Save(campaing *campaing.Campaing) error {
	c.campaings = append(c.campaings, *campaing)
	return nil
}

func (c *CampaingRepository) Get() []campaing.Campaing {
	return c.campaings
}
