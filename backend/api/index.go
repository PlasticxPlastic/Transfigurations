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
	config.AllowOrigins = []string{
		"http://localhost:5173",  // Vue.js dev server
		"https://transfiguration-frontend.vercel.app", // Production frontend
	}
	config.AllowCredentials = true
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