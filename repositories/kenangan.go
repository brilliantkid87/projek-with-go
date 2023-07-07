package repositories

import (
	"kenangan/model"

	"gorm.io/gorm"
)

type KenanganRepository interface {
	CreateKenangan(kenangan model.Kenangan) (model.Kenangan, error)
	GetKenanganByID(id int) (model.Kenangan, error)
	// GetAllKenangan() ([]model.Kenangan, error)
}

// connection
func RepositoryKenangan(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateKenangan(kenangan model.Kenangan) (model.Kenangan, error) {
	err := r.db.Create(&kenangan).Error

	return kenangan, err
}

func (r *repository) GetKenanganByID(id int) (model.Kenangan, error) {
	var kenangan model.Kenangan

	err := r.db.First(&kenangan, id).Error

	return kenangan, err
}
