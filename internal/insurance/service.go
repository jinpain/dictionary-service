package insurance

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetById(uuid *uuid.UUID) (*Model, error) {
	return nil, nil
}

func (s *Service) GetList(limit int, offset int) ([]*Model, error) {
	return nil, nil
}

func (s *Service) Create(model *Model) (*uuid.UUID, error) {
	return nil, nil
}

func (s *Service) Update(model *Model) error {
	return nil
}

func (s *Service) MarkDelete(uuid *uuid.UUID) error {
	return nil
}

func (s *Service) UnMarkDelete(uuid *uuid.UUID) error {
	return nil
}
