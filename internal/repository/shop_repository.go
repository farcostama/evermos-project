package repository

import (
	"evermos-project/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type ShopRepository interface {
	FindByUserID(userID uint) (*entity.Toko, error)
	FindByID(id uint) (*entity.Toko, error)
	Update(shop *entity.Toko) error
	FindAll(page, limit int, name string) ([]entity.Toko, int64, error) // Return data + total count
}

type shopRepository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &shopRepository{db}
}

func (r *shopRepository) FindByUserID(userID uint) (*entity.Toko, error) {
	var shop entity.Toko
	err := r.db.Where("user_id = ?", userID).First(&shop).Error
	return &shop, err
}

func (r *shopRepository) FindByID(id uint) (*entity.Toko, error) {
	var shop entity.Toko
	err := r.db.First(&shop, id).Error
	return &shop, err
}

func (r *shopRepository) Update(shop *entity.Toko) error {
	return r.db.Save(shop).Error
}

func (r *shopRepository) FindAll(page, limit int, name string) ([]entity.Toko, int64, error) {
	var shops []entity.Toko
	var total int64
	
	offset := (page - 1) * limit
	query := r.db.Model(&entity.Toko{})

	// Filtering
	if name != "" {
		query = query.Where("nama_toko LIKE ?", fmt.Sprintf("%%%s%%", name))
	}

	// Hitung total data untuk pagination
	query.Count(&total)

	// Ambil data dengan limit & offset (Pagination)
	err := query.Limit(limit).Offset(offset).Find(&shops).Error
	return shops, total, err
}