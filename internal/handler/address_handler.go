package handler

import (
	"evermos-project/internal/entity"
	"evermos-project/internal/service"
	"evermos-project/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AddressHandler struct {
	service service.AddressService
}

func NewAddressHandler(service service.AddressService) *AddressHandler {
	return &AddressHandler{service}
}

// GET /user/alamat (List)
func (h *AddressHandler) GetMyAddresses(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))

	addresses, err := h.service.GetMyAddresses(userID)
	if err != nil {
		return utils.ResponseJSON(c, 500, "Failed to GET data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to GET data", addresses, nil)
}

// GET /user/alamat/:id (Detail)
func (h *AddressHandler) GetAddressDetail(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id, _ := strconv.Atoi(c.Params("id"))

	address, err := h.service.GetAddressDetail(userID, uint(id))
	if err != nil {
		return utils.ResponseJSON(c, 404, "Failed to GET data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to GET data", address, nil)
}

// POST /user/alamat (Create)
func (h *AddressHandler) CreateAddress(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	var req entity.Alamat

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Request", nil, err.Error())
	}

	if err := h.service.CreateAddress(userID, &req); err != nil {
		return utils.ResponseJSON(c, 500, "Failed to POST data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to POST data", 1, nil)
}

// PUT /user/alamat/:id (Update)
func (h *AddressHandler) UpdateAddress(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id, _ := strconv.Atoi(c.Params("id"))
	var req entity.Alamat

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Request", nil, err.Error())
	}

	if err := h.service.UpdateAddress(userID, uint(id), &req); err != nil {
		return utils.ResponseJSON(c, 400, "Failed to UPDATE data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to UPDATE data", "Update succeed", nil)
}

// DELETE /user/alamat/:id (Delete)
func (h *AddressHandler) DeleteAddress(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.service.DeleteAddress(userID, uint(id)); err != nil {
		return utils.ResponseJSON(c, 400, "Failed to DELETE data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to DELETE data", "Delete succeed", nil)
}