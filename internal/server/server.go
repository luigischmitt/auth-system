package server

import (
	"auth-jwt/internal/database"
	"auth-jwt/internal/repositories"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	usersRepository repositories.UserRepositoryInterface
}

func NewServer() *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db := database.New()
	userRepository := repositories.NewUserRepository(db)

	server := &Server{
		usersRepository: userRepository,
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      server.RegisterRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	authHandler := NewAuthHandler(s.usersRepository)
	userHandler := NewUserHandler(s.usersRepository)

	// Setup routes
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/refresh", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true})
		})
		authRoutes.POST("/logout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true})
		})
	}

	protected := r.Group("/api")
	{
		protected.GET("/user", userHandler.GetUser)
	}

	return r
}
