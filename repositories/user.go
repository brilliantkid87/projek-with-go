package repositories

import (
	"kenangan/model"

	"gorm.io/gorm"
)

type UserRespository interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByID(id int) (model.User, error)
	// GetAllKenangan() ([]model.Kenangan, error)
}

// connection
func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) GetUserByID(id int) (model.User, error) {
	var user model.User

	err := r.db.Preload("Kenangan").First(&user, id).Error

	return user, err
}
