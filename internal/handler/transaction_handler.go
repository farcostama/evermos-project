package handler

import (
	"evermos-project/internal/service"
	"evermos-project/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type TransactionHandler struct {
	service service.TransactionService
}

func NewTransactionHandler(service service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service}
}

// POST /transaksi (Membuat transaksi baru)
func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	// Ambil ID Pembeli dari Token
	userID := uint(c.Locals("user_id").(float64))

	var req service.CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ResponseJSON(c, 400, "Invalid Request Format", nil, err.Error())
	}
    
	if err := h.service.CreateTransaction(userID, &req); err != nil {
		// Error dari service (misal: stok, alamat, harga)
		return utils.ResponseJSON(c, 400, "Transaction Failed", nil, err.Error())
	}

	return utils.ResponseJSON(c, 200, "Transaction Created Successfully", "Success", nil)
}
