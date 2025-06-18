package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var r *gin.Engine

func init() {
	r = gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true  // Allow all origins for testing
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}
	r.Use(cors.New(config))

	// Routes
	r.POST("/api/login", handleLogin)
}

// Handler is the main entry point for Vercel
func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
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

	// For testing, accept any username/password
	c.JSON(200, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id": 1,
			"username": loginData.Username,
		},
	})
} 