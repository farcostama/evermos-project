package repository

import (
	"evermos-project/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(page, limit int, name, categoryID, tokoID, minPrice, maxPrice string) ([]entity.Produk, int64, error)
	FindByID(id uint) (*entity.Produk, error)
	Create(product *entity.Produk) error
	Update(product *entity.Produk) error
	Delete(id uint) error
	
	// Khusus Foto Produk
	AddPhoto(photo *entity.FotoProduk) error
	DeletePhoto(photoID uint) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll(page, limit int, name, categoryID, tokoID, minPrice, maxPrice string) ([]entity.Produk, int64, error) {
	var products []entity.Produk
	var total int64
	
	offset := (page - 1) * limit
	
	// Preload Category dan Photos agar datanya lengkap saat diambil
	// Preload Toko agar info toko penjual juga muncul
	query := r.db.Model(&entity.Produk{}).Preload("Category").Preload("Photos").Preload("Toko")

	// --- FILTERING LOGIC ---
	if name != "" {
		query = query.Where("nama_produk LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	if tokoID != "" {
		query = query.Where("toko_id = ?", tokoID)
	}

	// --- PERBAIKAN LOGIC HARGA (CASTING STRING KE ANGKA) ---
	if minPrice != "" {
		// Mengubah string '75000' menjadi angka 75000 agar perbandingan akurat
		query = query.Where("CAST(harga_konsumen AS UNSIGNED) >= ?", minPrice)
	}
	if maxPrice != "" {
		query = query.Where("CAST(harga_konsumen AS UNSIGNED) <= ?", maxPrice)
	}
	// --------------------------------------------------------

	// Hitung Total Data (untuk Pagination)
	query.Count(&total)
	
	// Ambil Data dengan Limit & Offset
	err := query.Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}

func (r *productRepository) FindByID(id uint) (*entity.Produk, error) {
	var product entity.Produk
	// Ambil detail lengkap beserta Toko, Kategori, dan Foto-fotonya
	err := r.db.Preload("Toko").Preload("Category").Preload("Photos").First(&product, id).Error
	return &product, err
}

func (r *productRepository) Create(product *entity.Produk) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *entity.Produk) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	// Soft Delete (Karena di Entity ada DeletedAt)
	return r.db.Delete(&entity.Produk{}, id).Error
}

// --- FOTO ---
func (r *productRepository) AddPhoto(photo *entity.FotoProduk) error {
	return r.db.Create(photo).Error
}

func (r *productRepository) DeletePhoto(photoID uint) error {
	return r.db.Delete(&entity.FotoProduk{}, photoID).Error
}