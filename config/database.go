package config

import (
	"evermos-project/internal/entity"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] .env file not found. Using system environment variables.")
	}

	// Setup DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Connect ke Database
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// Panic akan menghentikan aplikasi dan mencetak error fatal
		log.Panicf("[FATAL] Failed to connect to database: %v", err)
	}

	log.Println("[INFO] Database connected successfully")

	// Auto Migrate
	log.Println("[INFO] Running database migration...")
	err = DB.AutoMigrate(
		&entity.User{},
		&entity.Alamat{},
		&entity.Toko{},
		&entity.Category{},
		&entity.Produk{},
		&entity.FotoProduk{},
		&entity.LogProduk{},
		&entity.Trx{},
		&entity.DetailTrx{},
	)
	if err != nil {
		log.Panicf("[FATAL] Database migration failed: %v", err)
	}
	
	log.Println("[INFO] Database migration completed successfully")
}