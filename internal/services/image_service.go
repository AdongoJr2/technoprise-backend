package services

import (
	"fmt"
	"github.com/AdongoJr2/technoprise-backend/internal/utils"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type ImageService struct {
	uploadDir string
	baseURL   string
}

func NewImageService(uploadDir, baseURL string) *ImageService {
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		return nil
	}

	return &ImageService{
		uploadDir: uploadDir,
		baseURL:   baseURL,
	}
}

func (s *ImageService) UploadImage(ctx echo.Context, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer src.Close()

	// Create a unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(s.uploadDir, filename)

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err = io.Copy(dst, src); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	// Return accessible URL
	publicHost := utils.GetPublicHost(ctx)
	return fmt.Sprintf("%s%s/%s", publicHost, s.baseURL, filename), nil
}
