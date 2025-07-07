package controllers

import (
	"github.com/AdongoJr2/technoprise-backend/internal/services"
	"github.com/AdongoJr2/technoprise-backend/internal/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// BlogPostHandler handles HTTP requests for blog posts.
type BlogPostHandler struct {
	service *services.BlogPostService
}

// NewBlogPostHandler creates a new BlogPostHandler.
func NewBlogPostHandler(service *services.BlogPostService) *BlogPostHandler {
	return &BlogPostHandler{service: service}
}

// CreateBlogPost handles the creation of a new blog post.
// POST /posts
func (h *BlogPostHandler) CreateBlogPost(c echo.Context) error {
	var input services.CreateBlogPostInput
	if err := c.Bind(&input); err != nil {
		return utils.NewHTTPError(http.StatusBadRequest, "Invalid request payload", err)
	}

	if input.Title == "" || input.Content == "" || input.Excerpt == "" {
		return utils.NewHTTPError(http.StatusBadRequest, "Title, Content and Excerpt are required", nil)
	}

	post, err := h.service.CreateBlogPost(c.Request().Context(), input)
	if err != nil {
		log.Printf("Handler error creating blog post: %v", err)
		return utils.NewHTTPError(http.StatusInternalServerError, "Failed to create blog post", err)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Blog post created successfully",
		"data":    post,
	})
}

// GetBlogPosts handles retrieving a list of blog posts with pagination and search.
// GET /posts?page=<int>&limit=<int>&search=<string>
func (h *BlogPostHandler) GetBlogPosts(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}

	searchTerm := c.QueryParam("search")

	paginatedPosts, err := h.service.GetBlogPosts(c.Request().Context(), page, limit, searchTerm)
	if err != nil {
		log.Printf("Handler error getting blog posts: %v", err)
		return utils.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve blog posts", err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "Blog posts retrieved successfully",
		"data":       paginatedPosts.Data,
		"pagination": paginatedPosts.Pagination,
	})
}

// GetBlogPostBySlug handles retrieving a single blog post by its slug.
// GET /posts/:slug
func (h *BlogPostHandler) GetBlogPostBySlug(c echo.Context) error {
	slug := c.Param("slug")
	if slug == "" {
		return utils.NewHTTPError(http.StatusBadRequest, "Slug is required", nil)
	}

	post, err := h.service.GetBlogPostBySlug(c.Request().Context(), slug)
	if err != nil {
		log.Printf("Handler error getting blog post by slug: %v", err)
		if err.Error() == "blog post with slug '"+slug+"' not found" {
			return utils.NewHTTPError(http.StatusNotFound, err.Error(), nil)
		}
		return utils.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve blog post", err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Blog post retrieved successfully",
		"data":    post,
	})
}
