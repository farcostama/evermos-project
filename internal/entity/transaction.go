package entity

import (
	"time"
)

// Trx (Transaction Header)
type Trx struct {
	ID               uint          `gorm:"primaryKey" json:"id"`
	UserID           uint          `json:"id_user"`
	User             User          `gorm:"foreignKey:UserID" json:"user"`

	AlamatPengirimanID uint          `json:"alamat_pengiriman_id"`
	AlamatPengiriman Alamat        `gorm:"foreignKey:AlamatPengirimanID" json:"alamat_pengiriman"`

	HargaTotal       string        `json:"harga_total"` 
	KodeInvoice      string        `json:"kode_invoice"`
	MethodBayar      string        `json:"method_bayar"`
	Status           string        `json:"status"` 

	DetailTrx        []DetailTrx   `gorm:"foreignKey:TrxID" json:"detail_trx"`
	
	CreatedAt        time.Time     `json:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at"`
}

// DetailTrx (Transaction Item Lines)
type DetailTrx struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TrxID         uint      `json:"id_trx"`

	LogProdukID   uint      `json:"id_log_produk"`
	LogProduk     LogProduk `gorm:"foreignKey:LogProdukID" json:"log_produk"` 

	TokoID        uint      `json:"id_toko"`
	Toko          Toko      `gorm:"foreignKey:TokoID" json:"toko"`

	Kuantitas     int       `json:"kuantitas"`
	HargaTotal    string    `json:"harga_total"` 

	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}