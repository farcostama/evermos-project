package middleware

import (
	"evermos-project/config"
	"evermos-project/internal/entity"
	"evermos-project/internal/utils"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware: Cek Token Valid
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return utils.ResponseJSON(c, 401, "Unauthorized", nil, "Missing Token")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return utils.ResponseJSON(c, 401, "Unauthorized", nil, "Invalid Token Format")
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return utils.ResponseJSON(c, 401, "Unauthorized", nil, "Invalid or Expired Token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return utils.ResponseJSON(c, 401, "Unauthorized", nil, "Invalid Token Claims")
	}

	// Simpan ID User ke Context
	c.Locals("user_id", claims["id"])
	
	return c.Next()
}

// AdminOnly: Cek Apakah User adalah Admin
func AdminOnly(c *fiber.Ctx) error {
	// Ambil ID dari Context (Hasil kerja AuthMiddleware)
	userID := c.Locals("user_id")
	if userID == nil {
		return utils.ResponseJSON(c, 401, "Unauthorized", nil, "User ID not found")
	}

	// Cek User di Database
	var user entity.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return utils.ResponseJSON(c, 404, "User Not Found", nil, "User tidak ditemukan")
	}

	// Validasi Status Admin
	if !user.IsAdmin {
		return utils.ResponseJSON(c, 403, "Forbidden", nil, "Akses Ditolak: Khusus Admin!")
	}

	return c.Next()
}