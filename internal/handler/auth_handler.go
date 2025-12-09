package handler

import (
	"evermos-project/internal/entity"
	"evermos-project/internal/service"
	"evermos-project/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

// POST /auth/register
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var user entity.User
	
	// Parse Body JSON
	if err := c.BodyParser(&user); err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Request", nil, err.Error())
	}

	// Panggil Service
	if err := h.service.Register(&user); err != nil {
		// Cek error duplicate entry (Email/NoTelp kembar)
		return utils.ResponseJSON(c, 400, "Failed to POST data", nil, []string{err.Error()})
	}

	return utils.ResponseJSON(c, 200, "Succeed to POST data", "Register Succeed", nil)
}

// POST /auth/login
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		NoTelp    string `json:"no_telp"`
		KataSandi string `json:"kata_sandi"`
	}

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Request", nil, err.Error())
	}

	token, user, err := h.service.Login(req.NoTelp, req.KataSandi)
	if err != nil {
		return utils.ResponseJSON(c, 401, "Failed to POST data", nil, []string{err.Error()})
	}

	// Format data response sesuai Postman
	data := fiber.Map{
		"nama":          user.Nama,
		"no_telp":       user.NoTelp,
		"tanggal_Lahir": user.TanggalLahir,
		"tentang":       user.Tentang,
		"pekerjaan":     user.Pekerjaan,
		"email":         user.Email,
		"id_provinsi":   user.IdProvinsi,
		"id_kota":       user.IdKota,
		"token":         token,
	}

	return utils.ResponseJSON(c, 200, "Succeed to POST data", data, nil)
}