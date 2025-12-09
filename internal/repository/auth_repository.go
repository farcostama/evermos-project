package repository

import (
	"evermos-project/internal/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RegisterUser(user *entity.User, toko *entity.Toko) error
	FindByPhone(phone string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

// RegisterUser: Simpan User dan Simpan Toko
// Menggunakan Transaction (Begin, Commit, Rollback) agar data aman.
func (r *authRepository) RegisterUser(user *entity.User, toko *entity.Toko) error {
	// Mulai Transaksi
	tx := r.db.Begin()

	// Simpan Data User
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback() 
		return err
	}

	// Set User ID untuk Toko (Relasi)
	toko.UserID = user.ID
	
	// Simpan Data Toko
	if err := tx.Create(toko).Error; err != nil {
		tx.Rollback() 
		return err
	}

	// Commit (Simpan permanen)
	return tx.Commit().Error
}

func (r *authRepository) FindByPhone(phone string) (*entity.User, error) {
	var user entity.User
	// Preload Toko agar data toko ikut terbawa saat login
	err := r.db.Where("no_telp = ?", phone).Preload("Toko").First(&user).Error
	return &user, err
}