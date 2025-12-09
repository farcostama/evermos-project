package service

import (
	"evermos-project/internal/entity"
	"evermos-project/internal/repository"
	"evermos-project/internal/utils"
)

type UserService interface {
	GetProfile(userID uint) (*entity.User, error)
	UpdateProfile(userID uint, req *entity.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetProfile(userID uint) (*entity.User, error) {
	return s.repo.FindByID(userID)
}

func (s *userService) UpdateProfile(userID uint, req *entity.User) error {
	// Ambil data lama dulu (biar data yg tidak diupdate tidak hilang)
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	//  Update field yang diizinkan saja
	if req.Nama != "" { user.Nama = req.Nama }
	if req.NoTelp != "" { user.NoTelp = req.NoTelp }
	if req.TanggalLahir != "" { user.TanggalLahir = req.TanggalLahir }
	if req.Pekerjaan != "" { user.Pekerjaan = req.Pekerjaan }
	if req.Email != "" { user.Email = req.Email }
	if req.IdProvinsi != "" { user.IdProvinsi = req.IdProvinsi }
	if req.IdKota != "" { user.IdKota = req.IdKota }
	if req.JenisKelamin != "" { user.JenisKelamin = req.JenisKelamin }
	if req.Tentang != "" { user.Tentang = req.Tentang }
	if req.KataSandi != "" {
		hashed, _ := utils.HashPassword(req.KataSandi)
		user.KataSandi = hashed
	}

	// Simpan perubahan
	return s.repo.Update(user)
}