package repository

import (
	"evermos-project/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*entity.User, error)
	Update(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	// ambil user, tapi passwordnya jangan diikutkan (aman)
	// Preload "Toko" agar data tokonya juga terlihat
	err := r.db.Preload("Toko").First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(user *entity.User) error {
	// Save akan mengupdate semua field yang ada di struct user
	return r.db.Save(user).Error
}