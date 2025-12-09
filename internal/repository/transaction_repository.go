package repository

import (
	"errors"
	"evermos-project/internal/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(trx *entity.Trx, details []entity.DetailTrx, logProduks []entity.LogProduk) error 
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

// CreateTransaction memastikan semua operasi DB berjalan serentak dan sukses (Atomic)
func (r *transactionRepository) CreateTransaction(trx *entity.Trx, details []entity.DetailTrx, logProduks []entity.LogProduk) error {
	
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Simpan Log Produk (Snapshot harga saat ini)
		if err := tx.Create(&logProduks).Error; err != nil {
			return err 
		}

		// Simpan Header Transaksi (Trx)
		if err := tx.Create(&trx).Error; err != nil {
			return err
		}

		// Mapping ID LogProduk dan ID Trx ke DetailTrx
		// Perbaikan: ID LogProduk dan ID Trx diisi di sini, setelah keduanya sukses disimpan
		for i := range details {
			details[i].TrxID = trx.ID // Hubungkan ke Header Transaksi yang baru dibuat
			details[i].LogProdukID = logProduks[i].ID // Ambil ID LogProduk yang baru disimpan
		}
		
		// Simpan Detail Transaksi (DetailTrx)
		if err := tx.Create(&details).Error; err != nil {
			return err
		}

		// Update Stok Produk (Decrement Stok)
		for i, detail := range details {
			// Kita ambil ProdukID dari LogProduk (snapshot)
			logProduk := logProduks[i] 

			// Update stok produk asli
			if err := tx.Model(&entity.Produk{}).
				Where("id = ?", logProduk.ProdukID).
				Update("stok", gorm.Expr("stok - ?", detail.Kuantitas)).
				Error; err != nil {
				return err
			}
			
			// Validasi Stok (Mencegah Stok Negatif)
			var updatedProduct entity.Produk
			tx.Select("stok").Where("id = ?", logProduk.ProdukID).First(&updatedProduct)
			if updatedProduct.Stok < 0 {
				return errors.New("stok tidak cukup, transaksi dibatalkan") 
			}
		}
		return nil
	})
}