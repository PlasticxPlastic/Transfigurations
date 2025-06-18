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
		"http://localhost:5173",
		"https://transfiguration1-qu1hr23f2-ohms-projects-4b3e1e96.vercel.app",
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
		"X-Requested-With",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
	}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * 60 * 60 // 12 hours
	r.Use(cors.New(config))

	// Add OPTIONS handler for preflight requests
	r.OPTIONS("/api/login", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "https://transfiguration1-qu1hr23f2-ohms-projects-4b3e1e96.vercel.app")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")
		c.Status(http.StatusOK)
	})

	// Routes
	r.POST("/api/login", handleLogin)
}

// Handler is the main entry point for Vercel
func Handler(w http.ResponseWriter, req *http.Request) {
	// Add CORS headers for all responses
	w.Header().Set("Access-Control-Allow-Origin", "https://transfiguration1-qu1hr23f2-ohms-projects-4b3e1e96.vercel.app")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	
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