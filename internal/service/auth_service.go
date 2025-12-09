package service

import (
	"errors"
	"evermos-project/internal/entity"
	"evermos-project/internal/repository"
	"evermos-project/internal/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Register(req *entity.User) error
	Login(phone, password string) (string, *entity.User, error)
}

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(req *entity.User) error {

	// Hash Password
	hashedPwd, err := utils.HashPassword(req.KataSandi)
	if err != nil {
		return err
	}
	req.KataSandi = hashedPwd

	// Siapkan Data Toko Default
	toko := entity.Toko{
		NamaToko: "Toko " + req.Nama, 
		UrlFoto:  "https://via.placeholder.com/150",
	}

	// Simpan ke Repo
	return s.repo.RegisterUser(req, &toko)
}

func (s *authService) Login(phone, password string) (string, *entity.User, error) {
	// Cari User by No Telp
	user, err := s.repo.FindByPhone(phone)
	if err != nil {
		return "", nil, errors.New("No Telp atau kata sandi salah")
	}

	// Cek Password
	if !utils.CheckPasswordHash(password, user.KataSandi) {
		return "", nil, errors.New("No Telp atau kata sandi salah")
	}

	// Generate JWT Token
	claims := jwt.MapClaims{
		"id":      user.ID,
		"no_telp": user.NoTelp,
		"role":    user.IsAdmin, 
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", nil, err
	}

	return t, user, nil
}