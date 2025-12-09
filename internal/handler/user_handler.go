package handler

import (
	"evermos-project/internal/entity"
	"evermos-project/internal/service"
	"evermos-project/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service}
}

// GET /user
func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	// Ambil ID dari Token (Hasil kerja Middleware)
	userID := c.Locals("user_id") 
	
	// Casting ke float64 (standar JWT numeric) lalu ke uint
	id := uint(userID.(float64))

	user, err := h.service.GetProfile(id)
	if err != nil {
		return utils.ResponseJSON(c, 404, "User not found", nil, err.Error())
	}

	return utils.ResponseJSON(c, 200, "Succeed to GET data", user, nil)
}

// PUT /user
func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	id := uint(userID.(float64))

	var req entity.User
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Request", nil, err.Error())
	}

	if err := h.service.UpdateProfile(id, &req); err != nil {
		return utils.ResponseJSON(c, 500, "Failed to UPDATE data", nil, err.Error())
	}

	return utils.ResponseJSON(c, 200, "Succeed to UPDATE data", "Update succeed", nil)
}