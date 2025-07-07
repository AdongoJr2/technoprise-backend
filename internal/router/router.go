package router

import (
	"github.com/AdongoJr2/technoprise-backend/internal/controllers"
	"github.com/AdongoJr2/technoprise-backend/internal/utils"
	"github.com/labstack/echo/v4"
)

// RegisterRoutes sets up all API routes for the application.
func RegisterRoutes(e *echo.Echo, blogPostController *controllers.BlogPostHandler) {
	// Set custom HTTP error handler
	e.HTTPErrorHandler = utils.CustomHTTPErrorHandler

	// Group API routes
	api := e.Group("/api/v1")

	// Blog Post Routes
	api.POST("/posts", blogPostController.CreateBlogPost)
	api.GET("/posts", blogPostController.GetBlogPosts)
	api.GET("/posts/:slug", blogPostController.GetBlogPostBySlug)

	// Health check route
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})
}
