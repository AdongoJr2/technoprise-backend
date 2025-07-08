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
	"strings"

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

	// Configure static file serving
	uploadsDir := "./uploads" // Local directory where images are stored
	publicPath := "/images"   // Public URL path

	// Prevent directory listing
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   uploadsDir,
		Browse: false,
		Index:  "",
	}))

	// Set proper cache headers
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Path(), "/images/") {
				c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
			}
			return next(c)
		}
	})

	e.Static(publicPath, uploadsDir)

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
	imageService := services.NewImageService(
		uploadsDir, // Local storage directory
		publicPath, // Base URL for accessing images
	)
	blogPostService := services.NewBlogPostService(client, imageService)
	blogPostController := controllers.NewBlogPostHandler(blogPostService, imageService)

	// Register routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Technoprise APIs")
	})
	router.RegisterRoutes(e, blogPostController)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.ServerPort)))
}
