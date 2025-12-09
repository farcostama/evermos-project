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
