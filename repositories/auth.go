package repositories

import (
	"kenangan/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user model.User) (model.User, error)
	Login(email string) (model.User, error)
	CheckAuth(ID int) (model.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) Login(email string) (model.User, error) {
	var user model.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) CheckAuth(ID int) (model.User, error) {
	var user model.User
	err := r.db.First(&user, ID).Error // add this code

	return user, err
}
