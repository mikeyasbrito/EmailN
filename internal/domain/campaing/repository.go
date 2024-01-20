package campaing

type Repository interface {
	Save(campaing *Campaing) error
}
