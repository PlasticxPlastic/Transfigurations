package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/transfiguration/internal/database"
)

func main() {
	// Initialize database connection
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Vue.js dev server
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// Routes
	r.POST("/api/login", handleLogin)

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func handleLogin(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// Get user from database
	user, err := database.GetUserByUsername(loginData.Username)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password (plain text comparison as requested)
	if user.Password != loginData.Password {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id": user.ID,
			"username": user.Username,
		},
	})
} 