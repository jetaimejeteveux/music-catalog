Here's an updated `README.md` that includes details about the **Tracks** and **Memberships** endpoints:  

---

# 🎵 Music Catalog

A **Music Catalog** API built with Golang, using **Gin**, **GORM**, and **JWT authentication**. This project follows **clean architecture** with a modular structure.

## 📂 Project Structure
```
music-catalog/
│── cmd/                     # Entry point of the application
│   └── main.go
│
│── internal/                 # Core business logic
│   ├── configs/              # Configuration management
│   ├── handlers/             # HTTP handlers (Controllers)
│   ├── middleware/           # Middleware (e.g., authentication, logging)
│   ├── models/               # Database models
│   ├── repository/           # Data access layer
│   ├── service/              # Business logic layer
│
│── pkg/                      # Reusable utilities
│   ├── httpclient/           # HTTP client helpers
│   ├── internalsql/          # Internal SQL helpers
│   ├── jwt/                  # JWT token management
│
│── config.yaml               # Application configuration file
│── docker-compose.yaml       # Docker setup for local development
│── go.mod                    # Go module file
│── go.sum                    # Dependency lock file
│── makefile                  # Automation scripts
│── .gitignore                 # Git ignore file
│── coverage.html             # Test coverage report
│── coverage.out              # Raw test coverage data
│── music-catalog.jpg         # Project logo or reference image
│── README.md                 # Project documentation
```

---

## 🚀 Features
✅ User Authentication (JWT)  
✅ CRUD Operations for Music Catalog  
✅ Track Search & Activity Tracking  
✅ Membership Signup & Login  
✅ PostgreSQL Database with GORM  
✅ Role-Based Access Control (RBAC)  
✅ Configurable via `config.yaml`  
✅ Structured Logging with **Zerolog**  
✅ Unit & Integration Testing using **Testify & SQLMock**  

---

## 🛠 Installation & Setup

### 1️⃣ Clone the Repository
```bash
git clone https://github.com/jetaimejeteveux/music-catalog.git
cd music-catalog
```

### 2️⃣ Install Dependencies
```bash
go mod tidy
```

### 3️⃣ Configure Environment
Edit `config.yaml` or use environment variables:
```yaml
server:
  port: 8080
database:
  host: localhost
  port: 5432
  user: your_user
  password: your_password
  dbname: music_catalog
jwt:
  secret: "your_secret_key"
```

### 4️⃣ Run with Docker
```bash
docker-compose up --build
```

### 5️⃣ Run Locally
```bash
go run cmd/main.go
```

---

## 🧪 Running Tests
Run all tests:
```bash
go test ./...
```
Generate test coverage report:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

---

## 📌 API Endpoints

### 🎼 Tracks API

| Method   | Endpoint                      | Description                        | Auth |
|----------|--------------------------------|------------------------------------|------|
| `GET`    | `/tracks/search?query=...`    | Search for tracks                  | ✅   |
| `POST`   | `/tracks/track-activity`      | Log user track activity            | ✅   |

### 🏆 Membership API

| Method   | Endpoint            | Description          | Auth |
|----------|---------------------|----------------------|------|
| `POST`   | `/memberships/signup` | Register a new user  | ❌   |
| `POST`   | `/memberships/login`  | Login and get token  | ❌   |

✅ = Requires JWT Token  

---

## 📜 License
This project is licensed under the MIT License.

---

## 📬 Contact
For any questions or issues, feel free to open an issue or reach out.
