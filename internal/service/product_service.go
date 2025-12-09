package service

import (
	"errors"
	"evermos-project/internal/entity"
	"evermos-project/internal/repository"
	"fmt"
	"strings"
	"time"
)

type ProductService interface {
	GetAllProducts(page, limit int, name, categoryID, tokoID, minPrice, maxPrice string) (map[string]interface{}, error)
	GetProductByID(id uint) (*entity.Produk, error)
	CreateProduct(userID uint, req *entity.Produk) error
	UpdateProduct(userID, productID uint, req *entity.Produk) error
	DeleteProduct(userID, productID uint) error
	AddProductPhoto(userID, productID uint, filename string) error
}

type productService struct {
	repo     repository.ProductRepository
	shopRepo repository.ShopRepository
}

// Kita butuh ShopRepo untuk cek apakah User ini punya Toko
func NewProductService(repo repository.ProductRepository, shopRepo repository.ShopRepository) ProductService {
	return &productService{repo, shopRepo}
}

// Helper: Ubah "Madu Asli" jadi "madu-asli-167823"
func makeSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	return fmt.Sprintf("%s-%d", slug, time.Now().Unix())
}

func (s *productService) GetAllProducts(page, limit int, name, categoryID, tokoID, minPrice, maxPrice string) (map[string]interface{}, error) {
	if page <= 0 { page = 1 }
	if limit <= 0 { limit = 10 }

	products, total, err := s.repo.FindAll(page, limit, name, categoryID, tokoID, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"page":  page,
		"limit": limit,
		"total": total,
		"data":  products,
	}, nil
}

func (s *productService) GetProductByID(id uint) (*entity.Produk, error) {
	return s.repo.FindByID(id)
}

func (s *productService) CreateProduct(userID uint, req *entity.Produk) error {
	// 1. Cari Toko milik User ini (Si Asep punya toko apa?)
	shop, err := s.shopRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("anda belum memiliki toko")
	}

	// 2. Set Data Otomatis
	req.TokoID = shop.ID // Produk otomatis masuk ke toko si Asep
	req.Slug = makeSlug(req.NamaProduk)

	return s.repo.Create(req)
}

func (s *productService) UpdateProduct(userID, productID uint, req *entity.Produk) error {
	// Cek Produk Ada?
	product, err := s.repo.FindByID(productID)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}

	// CEK KEPEMILIKAN (SECURITY)
	// Cari toko milik user yang login
	shop, err := s.shopRepo.FindByUserID(userID)
	
	// Jika toko tidak ditemukan, ATAU produk ini BUKAN milik toko user tersebut -> TOLAK
	if err != nil || product.TokoID != shop.ID {
		return errors.New("anda tidak berhak mengedit produk ini")
	}

	// Update
	if req.NamaProduk != "" {
		product.NamaProduk = req.NamaProduk
		product.Slug = makeSlug(req.NamaProduk)
	}
	if req.HargaReseller != "" { product.HargaReseller = req.HargaReseller }
	if req.HargaKonsumen != "" { product.HargaKonsumen = req.HargaKonsumen }
	if req.Stok != 0 { product.Stok = req.Stok }
	if req.Deskripsi != "" { product.Deskripsi = req.Deskripsi }
	if req.CategoryID != 0 { product.CategoryID = req.CategoryID }

	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(userID, productID uint) error {
	product, err := s.repo.FindByID(productID)
	if err != nil { return err }

	// Security Check
	shop, err := s.shopRepo.FindByUserID(userID)
	if err != nil || product.TokoID != shop.ID {
		return errors.New("anda tidak berhak menghapus produk ini")
	}

	return s.repo.Delete(productID)
}

func (s *productService) AddProductPhoto(userID, productID uint, filename string) error {
	// Security Check
	product, err := s.repo.FindByID(productID)
	if err != nil { return err }
	
	shop, err := s.shopRepo.FindByUserID(userID)
	if err != nil || product.TokoID != shop.ID {
		return errors.New("ilegal akses: ini bukan produk anda")
	}

	photo := entity.FotoProduk{
		ProdukID: productID,
		UrlFoto:  filename,
	}
	return s.repo.AddPhoto(&photo)
}