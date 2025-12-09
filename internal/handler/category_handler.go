package handler

import (
	"evermos-project/internal/entity"
	"evermos-project/internal/service"
	"evermos-project/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	data, err := h.service.GetAllCategories()
	if err != nil {
		return utils.ResponseJSON(c, 500, "Failed", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Success", data, nil)
}

func (h *CategoryHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	data, err := h.service.GetCategoryByID(uint(id))
	if err != nil {
		return utils.ResponseJSON(c, 404, "Not Found", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Success", data, nil)
}

func (h *CategoryHandler) Create(c *fiber.Ctx) error {
	var req entity.Category
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Bad Request", nil, err.Error())
	}

	if err := h.service.CreateCategory(&req); err != nil {
		return utils.ResponseJSON(c, 500, "Failed", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Category Created", "Success", nil)
}

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req entity.Category
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Bad Request", nil, err.Error())
	}

	if err := h.service.UpdateCategory(uint(id), &req); err != nil {
		return utils.ResponseJSON(c, 500, "Failed", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Category Updated", "Success", nil)
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.DeleteCategory(uint(id)); err != nil {
		return utils.ResponseJSON(c, 500, "Failed", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Category Deleted", "Success", nil)
}