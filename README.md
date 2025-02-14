# API Server in Go

This project is a simple REST API server built using the [Gin](https://github.com/gin-gonic/gin) framework in Go. The API provides basic CRUD operations for managing users, along with a dynamically generated random favicon.

## Features
- 🟢 **User Management**: Create, retrieve, and delete users.
- 🟢 **RESTful Endpoints**: Standard API conventions.
- 🟢 **Random Favicon**: Generates a new 16x16 favicon dynamically.
- 🟢 **Input Validation**: Ensures proper request handling.

## Installation
### Prerequisites
- Install [Go](https://golang.org/doc/install) (version 1.22 or later)

### Clone the Repository
```sh
git clone https://github.com/your-username/go-api-server.git
cd go-api-server
```

### Install Dependencies
```sh
go mod tidy
```

## Running the Server
```sh
go run main.go
```
The server starts at `http://localhost:8080`

## API Endpoints

### 🔹 Get All Users
**Request:**
```sh
GET /users
```
**Response:**
```json
[
  {"id": 1, "name": "Alice", "email": "alice@example.com"},
  {"id": 2, "name": "Bob", "email": "bob@example.com"}
]
```

### 🔹 Get a Specific User by ID
**Request:**
```sh
GET /users/2
```
**Response:**
```json
{"id": 2, "name": "Bob", "email": "bob@example.com"}
```

### 🔹 Create a New User
**Request:**
```sh
POST /users
Content-Type: application/json
{
  "name": "Charlie",
  "email": "charlie@example.com"
}
```
**Response:**
```json
{"id": 3, "name": "Charlie", "email": "charlie@example.com"}
```

### 🔹 Delete a User
**Request:**
```sh
DELETE /users/2
```
**Response:**
```json
{"message": "User deleted"}
```

### 🔹 Browser Hello Endpoint
**Request:**
```sh
GET /
```
**Response:**
```text
Hello, world!
```

### 🔹 Get Random Favicon
**Request:**
```sh
GET /favicon.ico
```
This generates a **16x16 random favicon** each time.

## Next Steps
🔹 **Enhancements**: Add persistent storage (PostgreSQL, MySQL).  
🔹 **Authentication**: Implement JWT-based authentication.  
🔹 **Deployment**: Deploy using Docker & Kubernetes.  

## License
MIT License © 2025 Your Name

