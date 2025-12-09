package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID            uint           `gorm:"primaryKey" json:"id"`
    Nama          string         `json:"nama"`
    KataSandi    string          `json:"kata_sandi"`
    NoTelp        string         `gorm:"unique" json:"no_telp"`
    TanggalLahir  string         `json:"tanggal_Lahir"`
    JenisKelamin  string         `json:"jenis_kelamin"`
    Tentang       string         `json:"tentang"`
    Pekerjaan     string         `json:"pekerjaan"`
    Email         string         `gorm:"unique" json:"email"`
    IdProvinsi    string         `json:"id_provinsi"`
    IdKota        string         `json:"id_kota"`
    IsAdmin       bool           `gorm:"default:false" json:"is_admin"`
    CreatedAt     time.Time      `json:"created_at"`
    UpdatedAt     time.Time      `json:"updated_at"`
    DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

    // Relasi
    Toko   Toko     `gorm:"foreignKey:UserID" json:"toko"`
    Alamat []Alamat `gorm:"foreignKey:UserID" json:"alamat"`
}