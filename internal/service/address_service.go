package service

import (
	"errors"
	"evermos-project/internal/entity"
	"evermos-project/internal/repository"
)

type AddressService interface {
	GetMyAddresses(userID uint) ([]entity.Alamat, error)
	GetAddressDetail(userID, addressID uint) (*entity.Alamat, error)
	CreateAddress(userID uint, req *entity.Alamat) error
	UpdateAddress(userID, addressID uint, req *entity.Alamat) error
	DeleteAddress(userID, addressID uint) error
}

type addressService struct {
	repo repository.AddressRepository
}

func NewAddressService(repo repository.AddressRepository) AddressService {
	return &addressService{repo}
}

func (s *addressService) GetMyAddresses(userID uint) ([]entity.Alamat, error) {
	return s.repo.FindByUserID(userID)
}

func (s *addressService) GetAddressDetail(userID, addressID uint) (*entity.Alamat, error) {
	address, err := s.repo.FindByID(addressID)
	if err != nil {
		return nil, err
	}
	// CEK KEPEMILIKAN: Apakah alamat ini milik user yang sedang login?
	if address.UserID != userID {
		return nil, errors.New("tidak berhak mengakses alamat ini")
	}
	return address, nil
}

func (s *addressService) CreateAddress(userID uint, req *entity.Alamat) error {
	req.UserID = userID 
	return s.repo.Create(req)
}

func (s *addressService) UpdateAddress(userID, addressID uint, req *entity.Alamat) error {
	// Cek dulu apakah alamatnya ada dan milik user ini
	address, err := s.GetAddressDetail(userID, addressID)
	if err != nil {
		return err
	}

	// Update data
	if req.JudulAlamat != "" { address.JudulAlamat = req.JudulAlamat }
	if req.NamaPenerima != "" { address.NamaPenerima = req.NamaPenerima }
	if req.NoTelp != "" { address.NoTelp = req.NoTelp }
	if req.DetailAlamat != "" { address.DetailAlamat = req.DetailAlamat }

	return s.repo.Update(address)
}

func (s *addressService) DeleteAddress(userID, addressID uint) error {
	// Cek kepemilikan dulu
	_, err := s.GetAddressDetail(userID, addressID)
	if err != nil {
		return err
	}
	// Hapus
	return s.repo.Delete(addressID)
}