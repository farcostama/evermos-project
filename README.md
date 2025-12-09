<div align="center">

# Evermos

### E-Commerce Simple Transaction Service

Build with **Go + Fiber** | Clean Architecture | Production Ready

</div>

---

## Overview

**Evermos-Project** adalah REST API service untuk transaksi penjualan e-commerce yang fokus pada **checkout & order management**. Service ini menyediakan fitur lengkap untuk user management, product catalog, dan transaksi dengan validasi stok real-time.

```
User Register -> Browse Products -> Add Address -> Checkout OK
```

---

## Features

| Category                 | Description                                           |
| ------------------------ | ----------------------------------------------------- |
| **User Management**      | Register, login, profile (JWT 72h)                    |
| **Product Catalog**      | CRUD + filtering (name, category, price) + pagination |
| **Address Management**   | Multiple address per user dengan ownership validation |
| **Checkout Transaction** | Multi-item transaction dengan stock validation        |
| **Category Management**  | Admin-only CRUD operations                            |
| **File Upload**          | Product photo storage                                 |
| **Clean Architecture**   | 4-layer pattern                                       |
| **Security**             | JWT, bcrypt, ownership validation                     |

---

## Quick Start

### Prerequisites

```bash
Go 1.21+  |  MySQL 8.0+  |  Git
```

### Installation

```bash
git clone <repository-url>
cd evermos-project
go mod download && go mod tidy

cp .env.example .env
# Edit: DATABASE_URL, JWT_SECRET

mysql -u root -p -e "CREATE DATABASE evermos CHARACTER SET utf8mb4;"
go run ./cmd/main.go
```

**Server:** http://localhost:8000

---

## Project Structure

```
evermos-project/
├── cmd/main.go                 Entry point & routing
├── config/database.go          Database configuration
├── internal/
│   ├── entity/                 Domain models
│   ├── repository/             Data access layer
│   ├── service/                Business logic
│   ├── handler/                HTTP endpoints
│   ├── middleware/             JWT & auth
│   └── utils/                  Helper functions
├── uploads/                    Product photos
└── .env.example
```

**Architecture:** Handler -> Service -> Repository -> Entity

---

## API Endpoints

### Public (No Auth)

```
POST   /api/v1/auth/register       Register & auto create shop
POST   /api/v1/auth/login          Login, get JWT token
GET    /api/v1/produk              List products (filter, pagination)
GET    /api/v1/produk/:id          Product detail
GET    /api/v1/toko                List shops
GET    /api/v1/category            List categories
```

### Protected (JWT Required)

```
GET    /api/v1/user                Get profile
PUT    /api/v1/user                Update profile
POST   /api/v1/user/alamat         Create address
GET    /api/v1/user/alamat         List addresses
PUT    /api/v1/user/alamat/:id     Update address
DELETE /api/v1/user/alamat/:id     Delete address
POST   /api/v1/produk              Create product
PUT    /api/v1/produk/:id          Update product
DELETE /api/v1/produk/:id          Delete product
POST   /api/v1/produk/foto         Upload photo
POST   /api/v1/transaksi           Create transaction (CHECKOUT)
```

### Admin (JWT + Admin Role)

```
POST   /api/v1/category            Create category
PUT    /api/v1/category/:id        Update category
DELETE /api/v1/category/:id        Delete category
```

---

## Usage Examples

### 1. Register User

```bash
curl -X POST http://localhost:8000/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "nama": "John Doe",
    "kata_sandi": "SecurePass123!",
    "no_telp": "08xxxxxxxxxx",
    "email": "john@example.com"
  }'
```

**Response (201):**

```json
{
  "data": {
    "id": 1,
    "nama": "John Doe",
    "email": "john@example.com"
  },
  "message": "Register berhasil! Toko otomatis dibuat."
}
```

### 2. Login

```bash
curl -X POST http://localhost:8000/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "no_telp": "08xxxxxxxxxx",
    "kata_sandi": "SecurePass123!"
  }'
```

**Response (200):**

```json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user_id": 1
  },
  "message": "Login berhasil"
}
```

### 3. Get Products (Filter & Pagination)

```bash
curl "http://localhost:8000/api/v1/produk?page=1&limit=10&name=madu&category_id=1&min_price=50000&max_price=500000"
```

### 4. Checkout (Create Transaction)

```bash
curl -X POST http://localhost:8000/api/v1/transaksi \
  -H "Authorization: Bearer <jwt-token>" \
  -H "Content-Type: application/json" \
  -d '{
    "alamat_id": 1,
    "items": [
      { "product_id": 5, "kuantitas": 2 },
      { "product_id": 8, "kuantitas": 1 }
    ]
  }'
```

**Response (201):**

```json
{
  "data": {
    "id": 1,
    "user_id": 1,
    "total_harga": 525000,
    "status": "pending",
    "items": [
      {
        "produk_id": 5,
        "kuantitas": 2,
        "harga_saat_beli": "175000"
      }
    ]
  },
  "message": "Transaksi berhasil dibuat"
}
```

---

## Tech Stack

| Layer              | Technology                  |
| ------------------ | --------------------------- |
| **Language**       | Go 1.21+                    |
| **Web Framework**  | Fiber v2                    |
| **Database**       | MySQL 8.0+ / PostgreSQL 12+ |
| **ORM**            | GORM                        |
| **Authentication** | JWT HS256 (72h)             |
| **Password**       | bcrypt (cost 10)            |
| **File Upload**    | multipart/form-data         |
| **Config**         | godotenv                    |

---

## Security

### Implementation

| Feature            | Technology                            |
| ------------------ | ------------------------------------- |
| **Authentication** | JWT HS256 Token (72h expiry)          |
| **Password**       | bcrypt hashing (cost 10)              |
| **Authorization**  | Ownership validation at service layer |
| **SQL Injection**  | GORM prepared statements              |
| **Constraints**    | Unique email & phone number           |
| **CORS**           | Fiber CORS middleware                 |

### Before Production

- [ ] Enable HTTPS/SSL
- [ ] JWT_SECRET: 32+ characters
- [ ] Database credentials in .env
- [ ] Add .env to .gitignore
- [ ] Setup error logging
- [ ] Enable database backups

---

## Configuration

Create `.env` file:

```env
DATABASE_URL=user:password@tcp(localhost:3306)/evermos?charset=utf8mb4
DB_DRIVER=mysql
PORT=8000
APP_NAME=evermos-service
JWT_SECRET=your-super-secret-key-minimum-32-characters
UPLOAD_PATH=./uploads
MAX_UPLOAD_SIZE=10485760
```

---

## Deployment

### Local Development

```bash
go run ./cmd/main.go
```

### Production Build

```bash
go build -o evermos ./cmd/main.go
./evermos
```

### Docker

```bash
docker build -t evermos:1.0 .
docker run -p 8000:8000 evermos:1.0
```

---

## Troubleshooting

| Problem                        | Solution                                          |
| ------------------------------ | ------------------------------------------------- |
| **Database connection failed** | Check DATABASE_URL & credentials in .env          |
| **Invalid token / 401 error**  | Verify: Authorization: Bearer <token>             |
| **Upload folder not found**    | Create ./uploads directory with write permissions |
| **Port already in use**        | Change PORT in .env                               |

---

## Requirements Checklist

- [x] User registration & login (JWT)
- [x] Password hashing (bcrypt)
- [x] Auto shop creation on register
- [x] Product CRUD + filtering + pagination
- [x] Address CRUD with ownership validation
- [x] Category management (admin-only)
- [x] File upload (product photos)
- [x] Transaction/Checkout with stock validation
- [x] Log product snapshots
- [x] Clean architecture (4-layer)
- [x] Standardized API responses
- [x] Error handling & status codes
- [x] Environment configuration
- [x] Database auto-migration

---

## Key Notes

- All business logic in service layer for testability
- Ownership validation at service level
- Atomic transactions for data consistency
- Auto-migration on application startup
- Standardized responses across endpoints

---

<div align="center">

### Ready for Production!

**[View API](#api-endpoints)** | **[Setup](#quick-start)** | **[Troubleshooting](#troubleshooting)**

Made with passion for Rakamin Task - Evermos

</div>
