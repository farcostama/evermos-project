package entity

import "time"

type Toko struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `json:"id_user"`
    NamaToko  string    `json:"nama_toko"`
    UrlFoto   string    `json:"url_foto"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}