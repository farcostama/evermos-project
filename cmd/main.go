package main

import (
	"evermos-project/config"
	"evermos-project/internal/handler"
	"evermos-project/internal/middleware"
	"evermos-project/internal/repository"
	"evermos-project/internal/service"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Konek Database
	config.ConnectDB()

	// Setup Dependency Injection
	
	// AUTH & USER
	authRepo := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// ADDRESS
	addressRepo := repository.NewAddressRepository(config.DB)
	addressService := service.NewAddressService(addressRepo)
	addressHandler := handler.NewAddressHandler(addressService)

	// SHOP
	shopRepo := repository.NewShopRepository(config.DB)
	shopService := service.NewShopService(shopRepo)
	shopHandler := handler.NewShopHandler(shopService)

	// CATEGORY
	categoryRepo := repository.NewCategoryRepository(config.DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService) 

	// PRODUCT
	productRepo := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepo, shopRepo) 
	productHandler := handler.NewProductHandler(productService)

	// TRANSACTION (FINAL)
	trxRepo := repository.NewTransactionRepository(config.DB)
	// Inject ProductRepo dan AddressRepo
	trxService := service.NewTransactionService(trxRepo, productRepo, addressRepo) 
	trxHandler := handler.NewTransactionHandler(trxService) 

	// Setup Fiber
	app := fiber.New()
	app.Use(cors.New())
	app.Static("/uploads", "./uploads") 

	// Routing Group
	api := app.Group("/api/v1")

	// --- PUBLIC ROUTES ---
	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)
	api.Get("/toko", shopHandler.GetAllShops)
	api.Get("/toko/:id", shopHandler.GetShopByID)
	api.Get("/category", categoryHandler.GetAll)
	api.Get("/category/:id", categoryHandler.GetByID)
	api.Get("/produk", productHandler.GetAllProducts)
	api.Get("/produk/:id", productHandler.GetProductByID)


	// --- PRIVATE ROUTES (Token Required) ---
	
	// Transaksi
	api.Post("/transaksi", middleware.AuthMiddleware, trxHandler.CreateTransaction) // POST /api/v1/transaksi

	// User & Alamat
	userAPI := api.Group("/user", middleware.AuthMiddleware)
	userAPI.Get("/", userHandler.GetProfile)
	userAPI.Put("/", userHandler.UpdateProfile)
	userAPI.Get("/alamat", addressHandler.GetMyAddresses)
	userAPI.Post("/alamat", addressHandler.CreateAddress)
	userAPI.Get("/alamat/:id", addressHandler.GetAddressDetail)
	userAPI.Put("/alamat/:id", addressHandler.UpdateAddress)
	userAPI.Delete("/alamat/:id", addressHandler.DeleteAddress)

	// Toko Private
	api.Get("/toko/my", middleware.AuthMiddleware, shopHandler.GetMyShop)
	api.Put("/toko/:id", middleware.AuthMiddleware, shopHandler.UpdateShop)
	
	// Produk Private
	api.Post("/produk", middleware.AuthMiddleware, productHandler.CreateProduct)
	api.Put("/produk/:id", middleware.AuthMiddleware, productHandler.UpdateProduct)
	api.Delete("/produk/:id", middleware.AuthMiddleware, productHandler.DeleteProduct)
	api.Post("/produk/foto", middleware.AuthMiddleware, productHandler.AddProductPhoto)

	// --- ADMIN ROUTES ---
	adminAPI := api.Group("/category", middleware.AuthMiddleware, middleware.AdminOnly)
	adminAPI.Post("/", categoryHandler.Create)
	adminAPI.Put("/:id", categoryHandler.Update)
	adminAPI.Delete("/:id", categoryHandler.Delete)

	// Jalankan Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Server running on port " + port)
	app.Listen(":" + port)
}