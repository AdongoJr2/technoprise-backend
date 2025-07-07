package main

import (
	"fmt"
	"github.com/AdongoJr2/technoprise-backend/internal/controllers"
	"github.com/AdongoJr2/technoprise-backend/internal/router"
	"github.com/AdongoJr2/technoprise-backend/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"

	"github.com/AdongoJr2/technoprise-backend/config"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Echo server
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allow all origins for development. Restrict in production.
	}))

	// Initialize database and Ent client
	client, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := client.Close(); err != nil {

			log.Printf("Failed to close database client: %v", err)
		}
	}()

	// Initialize services and handlers with the Ent client
	blogPostService := services.NewBlogPostService(client)
	blogPostController := controllers.NewBlogPostHandler(blogPostService)

	// Register routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Technoprise APIs")
	})
	router.RegisterRoutes(e, blogPostController)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.ServerPort)))
}
