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
	var model Model

	if err := s.db.Find(&model, uuid).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func (s *Service) GetList(filter *Filter) ([]*Model, error) {
	var models []*Model

	if err := s.db.Limit(filter.Limit).Offset(filter.Offset).Find(&models).Error; err != nil {
		return nil, err
	}

	return models, nil
}

func (s *Service) Create(model *Model) (*uuid.UUID, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	model.ID = uuid.String()

	if err := s.db.Create(model).Error; err != nil {
		return nil, err
	}

	return &uuid, nil
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
