package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"math/rand/v2"

	"github.com/gin-gonic/gin"
)

// User represents a user in the system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Log levels
const (
	LogLevelDebug = "debug"
	LogLevelInfo  = "info"
	LogLevelWarn  = "warn"
	LogLevelError = "error"
)

// Server struct to manage routes and the Gin engine
type Server struct {
	router *gin.Engine
	users  []User
}

// NewServer initializes a new API server
func NewServer() *Server {
	s := &Server{
		router: gin.Default(),
		users: []User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
			{ID: 2, Name: "Bob", Email: "bob@example.com"},
		},
	}
	s.setupRoutes()
	return s
}

// setupRoutes configures API endpoints
func (s *Server) setupRoutes() {
	s.router.GET("/users", s.getUsers)
	s.router.GET("/users/:id", s.getUsers)
	s.router.POST("/users", s.createUser)
	s.router.DELETE("/users/:id", s.deleteUser)

	//  Browser Hello
	s.router.GET("/", s.httpHello)

	// Favicon endpoint
	s.router.GET("/favicon.ico", s.faviconHandler)
}

// faviconHandler generates and returns a random favicon
func (s *Server) faviconHandler(c *gin.Context) {
	// Generate a 16x16 random image
	img := generateFavicon(16, 16)

	// Set response content type
	c.Writer.Header().Set("Content-Type", "image/png")

	// Encode image as PNG and write to response
	png.Encode(c.Writer, img)
}

// generateFavicon creates a simple random-colored favicon
func generateFavicon(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill with random colors using math/rand/v2
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r := uint8(rand.IntN(256)) // Random Red value
			g := uint8(rand.IntN(256)) // Random Green value
			b := uint8(rand.IntN(256)) // Random Blue value
			img.Set(x, y, color.RGBA{R: r, G: g, B: b, A: 255})
		}
	}
	return img
}

func (s *Server) httpHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Server Found": "Use API",
	})
}

// getUsers returns a list of all users
func (s *Server) getUsers(c *gin.Context) {
	// Check if an ID query parameter is provided
	idQuery := c.Query("id")
	if idQuery != "" {
		// Convert string to int
		id, err := strconv.Atoi(idQuery)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		// Search for the user by ID
		for _, user := range s.users {
			if user.ID == id {
				c.JSON(http.StatusOK, user) // Return the single user
				return
			}
		}

		// User not found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// If no ID is provided, return all users
	c.JSON(http.StatusOK, s.users)
}

// createUser adds a new user
func (s *Server) createUser(c *gin.Context) {
	var newUser User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	newUser.ID = len(s.users) + 1
	s.users = append(s.users, newUser)
	c.JSON(http.StatusCreated, newUser)
}

// deleteUser removes a user by ID
func (s *Server) deleteUser(c *gin.Context) {
	// Extract the ID from the request parameters
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Find and delete the user
	for i, user := range s.users {
		if user.ID == id {
			// Remove the user from the slice
			s.users = append(s.users[:i], s.users[i+1:]...)

			// Return success response
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}

	// If user not found, return 404
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

// Run starts the API server
func (s *Server) Run(addr string) {
	gin.SetMode(gin.ReleaseMode)
	s.router.Run(addr)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //Default to 8080 if not set
	}

	env := os.Getenv("APP_ENV")
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	if logLevel == "" {
		logLevel = LogLevelInfo // Default to "info"
	}

	// Configure log output based on log level
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // Adds timestamp & file info

	log.Printf("Starting API in %s mode on port %s with log level: %s", env, port, logLevel)

	// Example of logging at different levels
	switch logLevel {
	case LogLevelDebug:
		log.Println("[DEBUG] Debugging enabled")
		fallthrough
	case LogLevelInfo:
		log.Println("[INFO] Application is starting...")
	case LogLevelWarn:
		log.Println("[WARN] Warning level set")
	case LogLevelError:
		log.Println("[ERROR] Error logging only")
	default:
		log.Println("[INFO] Default log level: INFO")
	}

	server := NewServer()
	server.Run(":" + port)
}
