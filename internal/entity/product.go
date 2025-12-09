package entity

import (
	"time"

	"gorm.io/gorm"
)

type Produk struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	NamaProduk    string         `json:"nama_produk"`
	Slug          string         `json:"slug"` 
	HargaReseller string         `json:"harga_reseller"`
	HargaKonsumen string         `json:"harga_konsumen"`
	Stok          int            `json:"stok"`
	Deskripsi     string         `json:"deskripsi"`
	
	// RELASI KE TOKO
	TokoID        uint           `json:"toko_id"`
	Toko          Toko           `gorm:"foreignKey:TokoID" json:"toko"`
	
	// RELASI KE KATEGORI
	CategoryID    uint           `json:"category_id"`
	Category      Category       `gorm:"foreignKey:CategoryID" json:"category"`
	
	// RELASI KE FOTO (One to Many)
	Photos        []FotoProduk   `gorm:"foreignKey:ProdukID" json:"photos"`

	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type FotoProduk struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ProdukID  uint      `json:"product_id"`
	UrlFoto   string    `json:"url_foto"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}