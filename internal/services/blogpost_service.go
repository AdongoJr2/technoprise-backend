package services

import (
	"context"
	"fmt"
	"github.com/AdongoJr2/technoprise-backend/ent"
	"github.com/AdongoJr2/technoprise-backend/ent/blogpost"
	"github.com/AdongoJr2/technoprise-backend/internal/utils"
	"log"
	"math"
	"time"
)

// BlogPostService provides business logic for blog posts.
type BlogPostService struct {
	client *ent.Client
}

// NewBlogPostService creates a new BlogPostService.
func NewBlogPostService(client *ent.Client) *BlogPostService {
	return &BlogPostService{client: client}
}

// CreateBlogPostInput defines the input structure for creating a blog post.
type CreateBlogPostInput struct {
	Title   string `json:"title"`
	Excerpt string `json:"excerpt"`
	Content string `json:"content"`
}

// UpdateBlogPostInput defines the input structure for updating a blog post.
type UpdateBlogPostInput struct {
	Title   *string `json:"title,omitempty"`
	Excerpt *string `json:"excerpt,omitempty"`
	Content *string `json:"content,omitempty"`
}

// PaginatedBlogPosts holds blog posts and pagination metadata.
type PaginatedBlogPosts struct {
	Data       []*ent.BlogPost `json:"data"`
	Pagination PaginationMeta  `json:"pagination"`
}

// PaginationMeta holds pagination details.
type PaginationMeta struct {
	Total      int `json:"total"`
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"totalPages"`
}

// CreateBlogPost creates a new blog post in the database.
func (s *BlogPostService) CreateBlogPost(ctx context.Context, input CreateBlogPostInput) (*ent.BlogPost, error) {
	slug := utils.GenerateSlug(input.Title)

	existing, err := s.client.BlogPost.Query().Where(blogpost.SlugEQ(slug)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Printf("Error checking for existing slug: %v", err)
		return nil, fmt.Errorf("failed to check for existing slug: %w", err)
	}
	if existing != nil {
		slug = fmt.Sprintf("%s-%d", slug, time.Now().Unix())
	}

	post, err := s.client.BlogPost.
		Create().
		SetTitle(input.Title).
		SetSlug(slug).
		SetContent(input.Content).
		SetExcerpt(input.Excerpt).
		Save(ctx)
	if err != nil {
		log.Printf("Error creating blog post: %v", err)
		return nil, fmt.Errorf("failed to create blog post: %w", err)
	}
	return post, nil
}

// GetBlogPosts retrieves a list of blog posts with pagination and search.
func (s *BlogPostService) GetBlogPosts(ctx context.Context, page, limit int, searchTerm string) (*PaginatedBlogPosts, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10 // Default limit
	}

	query := s.client.BlogPost.Query()

	// Apply search filter if searchTerm is provided
	if searchTerm != "" {
		query = query.Where(
			blogpost.Or(
				blogpost.TitleContainsFold(searchTerm),
				blogpost.ContentContainsFold(searchTerm),
				blogpost.ExcerptContainsFold(searchTerm),
			),
		)
	}

	// Get total count for pagination
	total, err := query.Count(ctx)
	if err != nil {
		log.Printf("Error counting blog posts: %v", err)
		return nil, fmt.Errorf("failed to count blog posts: %w", err)
	}

	// Calculate pagination details
	offset := (page - 1) * limit
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	// Fetch posts with pagination, ordered by creation date descending
	posts, err := query.
		Select(
			blogpost.FieldID,
			blogpost.FieldTitle,
			blogpost.FieldSlug,
			blogpost.FieldExcerpt,
			blogpost.FieldCreateTime,
			blogpost.FieldUpdateTime,
		).
		Order(ent.Desc(blogpost.FieldCreateTime)).
		Offset(offset).
		Limit(limit).
		All(ctx)
	if err != nil {
		log.Printf("Error fetching paginated blog posts: %v", err)
		return nil, fmt.Errorf("failed to fetch blog posts: %w", err)
	}

	return &PaginatedBlogPosts{
		Data: posts,
		Pagination: PaginationMeta{
			Total:      total,
			Page:       page,
			Limit:      limit,
			TotalPages: totalPages,
		},
	}, nil
}

// GetBlogPostBySlug retrieves a single blog post by its slug.
func (s *BlogPostService) GetBlogPostBySlug(ctx context.Context, slug string) (*ent.BlogPost, error) {
	post, err := s.client.BlogPost.Query().Where(blogpost.SlugEQ(slug)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, fmt.Errorf("blog post with slug '%s' not found", slug)
		}
		log.Printf("Error fetching blog post by slug '%s': %v", slug, err)
		return nil, fmt.Errorf("failed to retrieve blog post: %w", err)
	}
	return post, nil
}
