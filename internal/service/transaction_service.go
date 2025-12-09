package service

import (
	"errors"
	"evermos-project/internal/entity"
	"evermos-project/internal/repository"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// ItemRequest adalah struktur yang diterima dari body JSON user
type ItemRequest struct {
	ProdukID  uint `json:"product_id" validate:"required"`
	Kuantitas int  `json:"quantity" validate:"required,gt=0"`
}

// CreateTransactionRequest adalah body utama untuk membuat transaksi
type CreateTransactionRequest struct {
	AlamatPengirimanID uint          `json:"address_id" validate:"required"`
	MethodBayar        string        `json:"payment_method" validate:"required"`
	Items              []ItemRequest `json:"items" validate:"required,dive"`
}

// ----------------------------------------------------------------------

type TransactionService interface {
	CreateTransaction(userID uint, req *CreateTransactionRequest) error
}

type transactionService struct {
	trxRepo     repository.TransactionRepository
	productRepo repository.ProductRepository
	addressRepo repository.AddressRepository
}

func NewTransactionService(
	trxRepo repository.TransactionRepository,
	productRepo repository.ProductRepository,
	addressRepo repository.AddressRepository,
) TransactionService {
	return &transactionService{trxRepo, productRepo, addressRepo}
}

func (s *transactionService) CreateTransaction(userID uint, req *CreateTransactionRequest) error {
	// Validasi Alamat (Pastikan alamat itu milik user yang sedang login)
	alamat, err := s.addressRepo.FindByID(req.AlamatPengirimanID)
	if err != nil || alamat.UserID != userID {
		return fmt.Errorf("alamat pengiriman tidak valid atau bukan milik anda")
	}

	// Persiapan data yang akan disimpan secara Atomic
	var totalHargaSemuaItem float64
	logProduks := []entity.LogProduk{}
	detailTrxs := []entity.DetailTrx{}

	// Map untuk menyimpan data produk yang sudah diambil dari DB (agar tidak redundant)
	productMap := make(map[uint]*entity.Produk)

	// Validasi Setiap Item, Buat Log Produk, Hitung Total
	for _, item := range req.Items {
		
		// Ambil Produk dari DB (dengan preload Toko & Kategori untuk Log)
		produk, ok := productMap[item.ProdukID]
		if !ok {
			// Jika belum ada di map, ambil dari DB
			produk, err = s.productRepo.FindByID(item.ProdukID)
			if err != nil {
				return fmt.Errorf("produk dengan ID %d tidak ditemukan", item.ProdukID)
			}
			productMap[item.ProdukID] = produk 
		}
		
		// Validasi Stok
		if produk.Stok < item.Kuantitas {
			return fmt.Errorf("stok produk %s tidak cukup. Sisa: %d", produk.NamaProduk, produk.Stok)
		}
		
		// Hitung Harga & Konversi
		hargaKonsumen, err := strconv.ParseFloat(produk.HargaKonsumen, 64)
		if err != nil {
			return errors.New("kesalahan konversi harga produk")
		}
		hargaItemTotal := hargaKonsumen * float64(item.Kuantitas)
		totalHargaSemuaItem += hargaItemTotal

		// Buat LogProduk (SNAPSHOT)
		log := entity.LogProduk{
			ProdukID:      produk.ID,
			NamaProduk:    produk.NamaProduk,
			Slug:          produk.Slug,
			HargaReseller: produk.HargaReseller,
			HargaKonsumen: produk.HargaKonsumen,
			Deskripsi:     produk.Deskripsi,
			TokoID:        produk.TokoID,
			CategoryID:    produk.CategoryID,
		}
		logProduks = append(logProduks, log)

		// Buat DetailTrx
		detailTrxs = append(detailTrxs, entity.DetailTrx{
			LogProdukID:   log.ID, 
			TokoID:        produk.TokoID,
			Kuantitas:     item.Kuantitas,
			HargaTotal:    fmt.Sprintf("%.2f", hargaItemTotal),
		})
	}
	
	// Buat Header Transaksi (Trx)
	// Catatan: KodeInvoice harus unik (gunakan timestamp + random)
	rand.Seed(time.Now().UnixNano())
	kodeInvoice := fmt.Sprintf("INV-%d-%d", time.Now().Unix(), rand.Intn(1000))

	trxHeader := entity.Trx{
		UserID:           userID,
		AlamatPengirimanID: req.AlamatPengirimanID,
		HargaTotal:       fmt.Sprintf("%.2f", totalHargaSemuaItem),
		KodeInvoice:      kodeInvoice,
		MethodBayar:      req.MethodBayar,
		Status:           "PENDING", 
	}

	// Eksekusi Atomic Transaction
	// Repo akan: [1. Simpan LogProduk] -> [2. Simpan Trx] -> [3. Simpan DetailTrx] -> [4. Kurangi Stok]
	if err := s.trxRepo.CreateTransaction(&trxHeader, detailTrxs, logProduks); err != nil {
		return err
	}

	return nil
}