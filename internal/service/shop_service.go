package service

import (
	"errors"
	"evermos-project/internal/entity"
	"evermos-project/internal/repository"
)

type ShopService interface {
	GetMyShop(userID uint) (*entity.Toko, error)
	GetShopByID(id uint) (*entity.Toko, error)
	UpdateShop(userID, shopID uint, req *entity.Toko) error
	GetAllShops(page, limit int, name string) (map[string]interface{}, error)
}

type shopService struct {
	repo repository.ShopRepository
}

func NewShopService(repo repository.ShopRepository) ShopService {
	return &shopService{repo}
}

func (s *shopService) GetMyShop(userID uint) (*entity.Toko, error) {
	return s.repo.FindByUserID(userID)
}

func (s *shopService) GetShopByID(id uint) (*entity.Toko, error) {
	return s.repo.FindByID(id)
}

func (s *shopService) UpdateShop(userID, shopID uint, req *entity.Toko) error {
	// Ambil data toko lama
	shop, err := s.repo.FindByID(shopID)
	if err != nil {
		return errors.New("Toko tidak ditemukan")
	}

	// CEK KEPEMILIKAN (Ketentuan 13)
	// Jika ID User yang login beda dengan pemilik toko -> TOLAK
	if shop.UserID != userID {
		return errors.New("anda tidak berhak mengelola toko ini")
	}

	// Update data
	if req.NamaToko != "" { shop.NamaToko = req.NamaToko }
	if req.UrlFoto != "" { shop.UrlFoto = req.UrlFoto }

	return s.repo.Update(shop)
}

func (s *shopService) GetAllShops(page, limit int, name string) (map[string]interface{}, error) {
	// Default pagination
	if page <= 0 { page = 1 }
	if limit <= 0 { limit = 10 }

	shops, total, err := s.repo.FindAll(page, limit, name)
	if err != nil {
		return nil, err
	}

	// Format Pagination sesuai Postman
	return map[string]interface{}{
		"page":  page,
		"limit": limit,
		"total": total,
		"data":  shops,
	}, nil
}