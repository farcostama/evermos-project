package handler

import (
	"evermos-project/internal/entity"
	"evermos-project/internal/service"
	"evermos-project/internal/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ShopHandler struct {
	service service.ShopService
}

func NewShopHandler(service service.ShopService) *ShopHandler {
	return &ShopHandler{service}
}

// GET /toko/my
func (h *ShopHandler) GetMyShop(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	shop, err := h.service.GetMyShop(userID)
	if err != nil {
		return utils.ResponseJSON(c, 404, "Toko not found", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to GET data", shop, nil)
}

// GET /toko (Pagination & Search)
func (h *ShopHandler) GetAllShops(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	name := c.Query("nama")

	data, err := h.service.GetAllShops(page, limit, name)
	if err != nil {
		return utils.ResponseJSON(c, 500, "Failed to GET data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to GET data", data, nil)
}

// GET /toko/:id
func (h *ShopHandler) GetShopByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	shop, err := h.service.GetShopByID(uint(id))
	if err != nil {
		return utils.ResponseJSON(c, 404, "Failed to GET data", nil, "Toko tidak ditemukan")
	}
	return utils.ResponseJSON(c, 200, "Succeed to GET data", shop, nil)
}

// PUT /toko/:id (Update + Upload Foto)
func (h *ShopHandler) UpdateShop(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	shopID, _ := strconv.Atoi(c.Params("id"))

	// Ambil Nama Toko dari Form-Data
	namaToko := c.FormValue("nama_toko")
	
	// Handle Upload File 
	file, err := c.FormFile("photo")
	var filename string

	if err == nil {
		// Jika ada file yang diupload, simpan ke folder "uploads"
		// Nama file buat unik Timestamp agar tidak bentrok
		filename = fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		
		// Simpan file
		if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
			return utils.ResponseJSON(c, 500, "Failed to save file", nil, err.Error())
		}
	}

	// Masukkan ke Struct Request
	req := entity.Toko{
		NamaToko: namaToko,
		UrlFoto:  filename, 
	}

	// Panggil Service
	if err := h.service.UpdateShop(userID, uint(shopID), &req); err != nil {
		return utils.ResponseJSON(c, 400, "Failed to UPDATE data", nil, err.Error())
	}

	return utils.ResponseJSON(c, 200, "Succeed to UPDATE data", "Update toko succeed", nil)
}