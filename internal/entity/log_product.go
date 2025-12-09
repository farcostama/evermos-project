package entity

import (
	"time"
)

type LogProduk struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	ProdukID      uint      `json:"id_produk"` // ID referensi ke produk asli (history)
	NamaProduk    string    `json:"nama_produk"`
	Slug          string    `json:"slug"`
	HargaReseller string    `json:"harga_reseller"`
	HargaKonsumen string    `json:"harga_konsumen"`
	Deskripsi     string    `json:"deskripsi"`
	
	TokoID        uint      `json:"id_toko"`
	Toko          Toko      `gorm:"foreignKey:TokoID" json:"toko"`

	CategoryID    uint      `json:"id_category"`
	Category      Category  `gorm:"foreignKey:CategoryID" json:"category"`

	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}