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

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

// GET /produk (Public - Search & Filter)
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	
	name := c.Query("nama_produk")
	categoryID := c.Query("category_id")
	tokoID := c.Query("toko_id")
	minPrice := c.Query("harga_min")
	maxPrice := c.Query("harga_max")

	data, err := h.service.GetAllProducts(page, limit, name, categoryID, tokoID, minPrice, maxPrice)
	if err != nil {
		return utils.ResponseJSON(c, 500, "Failed to GET data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to GET data", data, nil)
}

// GET /produk/:id (Public - Detail)
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		return utils.ResponseJSON(c, 404, "Product not found", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to GET data", product, nil)
}

// POST /produk (Private - Tambah Produk Data Saja)
func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	var req entity.Produk

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Request", nil, err.Error())
	}

	if err := h.service.CreateProduct(userID, &req); err != nil {
		return utils.ResponseJSON(c, 400, "Failed to POST data", nil, err.Error())
	}
	
	// kembalikan Data Produk (terutama ID-nya) agar User tau ID produk barunya
	return utils.ResponseJSON(c, 200, "Succeed to POST data", req, nil)
}

// PUT /produk/:id (Private - Edit Produk)
func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id, _ := strconv.Atoi(c.Params("id"))
	var req entity.Produk

	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Request", nil, err.Error())
	}

	if err := h.service.UpdateProduct(userID, uint(id), &req); err != nil {
		return utils.ResponseJSON(c, 400, "Failed to UPDATE data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to UPDATE data", "Product updated successfully", nil)
}

// DELETE /produk/:id (Private - Hapus Produk)
func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.service.DeleteProduct(userID, uint(id)); err != nil {
		return utils.ResponseJSON(c, 400, "Failed to DELETE data", nil, err.Error())
	}
	return utils.ResponseJSON(c, 200, "Succeed to DELETE data", "Product deleted successfully", nil)
}

// POST /produk/foto (Private - Upload Foto Produk)
// Menggunakan Form-Data: key "photo" (file) dan "product_id" (text)
func (h *ProductHandler) AddProductPhoto(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	
	// Ambil Product ID dari Form Data
	productIDStr := c.FormValue("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Product ID", nil, "product_id is required")
	}

	// Ambil File Foto
	file, err := c.FormFile("photo")
	if err != nil {
		return utils.ResponseJSON(c, 400, "Missing Photo", nil, err.Error())
	}

	// Simpan File ke Server
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	if err := c.SaveFile(file, "./uploads/"+filename); err != nil {
		return utils.ResponseJSON(c, 500, "Failed to save file", nil, err.Error())
	}

	// Panggil Service untuk simpan ke Database
	if err := h.service.AddProductPhoto(userID, uint(productID), filename); err != nil {
		// Kalau gagal simpan ke DB, hapus filenya
		// os.Remove("./uploads/"+filename) 
		return utils.ResponseJSON(c, 400, "Failed to Add Photo", nil, err.Error())
	}

	return utils.ResponseJSON(c, 200, "Succeed to POST data", "Photo uploaded successfully", nil)
}