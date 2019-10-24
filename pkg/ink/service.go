package ink

import "fmt"

type Repository interface {
	AddInk(Ink) error
	GetInk(int64) (*Ink, error)
	GetAllInks() (*[]Ink, error)
}

type Service interface {
	AddInk(Ink) error
	GetInk(int64) (*Ink, error)
	GetAllInks() (*[]Ink, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) AddInk(ink Ink) error {
	if ink.ColorFamily == 0 {
		return fmt.Errorf("colorFamily is required")
	}
	if ink.Name == "" {
		return fmt.Errorf("namme is required")
	}
	return s.repo.AddInk(ink)
}

func (s *service) GetInk(id int64) (*Ink, error) {
	return s.repo.GetInk(id)
}

func (s *service) GetAllInks() (*[]Ink, error) {
	return s.repo.GetAllInks()
}
