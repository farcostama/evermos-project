package repository

import (
	"evermos-project/internal/entity"

	"gorm.io/gorm"
)

type AddressRepository interface {
	FindByUserID(userID uint) ([]entity.Alamat, error)
	FindByID(id uint) (*entity.Alamat, error)
	Create(address *entity.Alamat) error
	Update(address *entity.Alamat) error
	Delete(id uint) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) FindByUserID(userID uint) ([]entity.Alamat, error) {
	var addresses []entity.Alamat
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *addressRepository) FindByID(id uint) (*entity.Alamat, error) {
	var address entity.Alamat
	err := r.db.First(&address, id).Error
	return &address, err
}

func (r *addressRepository) Create(address *entity.Alamat) error {
	return r.db.Create(address).Error
}

func (r *addressRepository) Update(address *entity.Alamat) error {
	return r.db.Save(address).Error
}

func (r *addressRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Alamat{}, id).Error
}