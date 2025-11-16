# Go API Starter (Echo + GORM + Postgres)

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/AllForOne7/go-api-starter/ci.yml)](https://github.com/AllForOne7/go-api-starter/actions)

A clean and robust starter template for building RESTful APIs in Go. This project demonstrates a full CRUD API for managing messages, using a modern and scalable tech stack.

## 📸 Screenshots

![API Demo](https://via.placeholder.com/800x400?text=API+Demo+Screenshot)
*Example of API response in Postman or similar tool.*

## 📋 Table of Contents

- [Features](#-features)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Configuration](#-configuration)
- [Usage](#-usage)
- [API Endpoints](#-api-endpoints)
- [Testing](#-testing)
- [Contributing](#-contributing)
- [Roadmap](#-roadmap)
- [License](#-license)
- [Authors](#-authors)

## ✨ Features

- **Go (Golang)**: Modern, high-performance language for backend development.
- **Echo Framework**: High-performance, minimalist web framework for Go.
- **PostgreSQL**: Powerful, open-source relational database.
- **GORM ORM**: Developer-friendly ORM with auto-migration support.
- **Clean Architecture**: Dependency Injection via Handler struct, no global variables.
- **Proper Error Handling**: Distinguishes between client (4xx) and server (5xx) errors.
- **Full CRUD Operations**: Create, Read, Update, Delete for Message model.
- **JSON API**: RESTful endpoints with JSON request/response.
- **Validation**: Basic input validation for data integrity.

## 🛠 Prerequisites

- [Go](https://golang.org/doc/install) (version 1.25+)
- [PostgreSQL](https://www.postgresql.org/download/) (running locally or via Docker)
- [Git](https://git-scm.com/downloads)

## 🚀 Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/AllForOne7/go-api-starter.git
   cd go-api-starter
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database:**
   - Ensure PostgreSQL is running.
   - Create a database named `postgres` (or update DSN in `main.go`).
   - Default DSN: `host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable`

4. **Run the application:**
   ```bash
   go run main.go handler.go model.go
   ```

   The server will start on `http://localhost:8080`.

## ⚙️ Configuration

The database connection is configured in `main.go`. Update the DSN string for your environment:

```go
dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
```

For production, consider using environment variables:

```go
import "os"
dsn := os.Getenv("DATABASE_URL")
```

## 📖 Usage

Once the server is running, you can interact with the API using tools like `curl`, Postman, or any HTTP client.

### Examples

**Get all messages:**
```bash
curl http://localhost:8080/messages
```

**Create a new message:**
```bash
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"text": "Hello, World!"}'
```

**Update a message:**
```bash
curl -X PATCH http://localhost:8080/messages/1 \
  -H "Content-Type: application/json" \
  -d '{"text": "Updated message"}'
```

**Delete a message:**
```bash
curl -X DELETE http://localhost:8080/messages/1
```

## 🔌 API Endpoints

| Method | Endpoint          | Description          |
|--------|-------------------|----------------------|
| GET    | `/messages`      | Get all messages    |
| POST   | `/messages`      | Create a new message |
| PATCH  | `/messages/:id`  | Update a message by ID |
| DELETE | `/messages/:id`  | Delete a message by ID |

### Response Format

Success responses return JSON data. Error responses follow this structure:

```json
{
  "status": "Error",
  "message": "Description of the error"
}
```

## 🧪 Testing

To run tests (if implemented):

```bash
go test ./...
```

Currently, no tests are included. Consider adding unit tests for handlers and integration tests for API endpoints.

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a feature branch: `git checkout -b feature/your-feature`.
3. Commit your changes: `git commit -m 'Add some feature'`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Open a Pull Request.

### Code Style

- Follow Go conventions (use `gofmt` and `go vet`).
- Write clear, concise commit messages.
- Add tests for new features.

## 🗺 Roadmap

- [ ] Add authentication (JWT)
- [ ] Implement logging middleware
- [ ] Add Docker support
- [ ] Write comprehensive tests
- [ ] Add API documentation (Swagger)
- [ ] Support for multiple databases

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

- **AllForOne7** - *Initial work* - [GitHub](https://github.com/AllForOne7)

See also the list of [contributors](https://github.com/AllForOne7/go-api-starter/contributors) who participated in this project.