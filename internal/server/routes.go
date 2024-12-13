package server

import (
	"go-auth/internal/auth"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	userRepo := user.NewUserRepository(s.db.GetDB())
	authRepo := auth.NewAuthRepository(s.db.GetDB())

	userSvc := user.NewUserService(userRepo)
	authSvc := auth.NewAuthService(authRepo)

	userHandler := user.NewUserHandler(userSvc)
	authHandler := auth.NewAuthHandler(authSvc)

	r.GET("/", s.HelloWorldHandler)

	r.GET("/health", s.healthHandler)

	apiGroup := r.Group("/api")

	// Auth routes
	authGroup := apiGroup.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}

	// User routes
	userGroup := apiGroup.Group("/user")
	{
		userGroup.GET("/", userHandler.GetUser)
		userGroup.GET("/all", middleware.JWTMiddleware(authSvc), userHandler.GetAllUsers)
	}

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, s.db.Health())
}
