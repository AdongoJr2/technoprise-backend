package utils

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

// HTTPError represents a custom HTTP error response.
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

// Error implements the error interface for HTTPError.
func (e *HTTPError) Error() string {
	if e.Details != "" {
		return e.Message + ": " + e.Details
	}
	return e.Message
}

// NewHTTPError creates a new HTTPError instance.
func NewHTTPError(code int, message string, err error) *HTTPError {
	httpErr := &HTTPError{
		Code:    code,
		Message: message,
	}
	if err != nil {
		httpErr.Details = err.Error()
	}
	return httpErr
}

// CustomHTTPErrorHandler handles custom HTTP errors.
func CustomHTTPErrorHandler(err error, c echo.Context) {
	var report *echo.HTTPError
	ok := errors.As(err, &report)
	if ok {
		// Handle Echo's built-in HTTP errors
		c.Logger().Error(report)
		err := c.JSON(report.Code, HTTPError{
			Code:    report.Code,
			Message: report.Message.(string),
		})
		if err != nil {
			return
		}
		return
	}

	// Handle custom HTTPError
	var customErr *HTTPError
	ok = errors.As(err, &customErr)
	if ok {
		c.Logger().Error(customErr)
		err := c.JSON(customErr.Code, customErr)
		if err != nil {
			return
		}
		return
	}

	// Default error handling for any other unhandled errors
	c.Logger().Error(err)
	err = c.JSON(http.StatusInternalServerError, HTTPError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Details: err.Error(),
	})
	if err != nil {
		return
	}
}

// GenerateSlug converts a string to a URL-friendly slug.
func GenerateSlug(s string) string {
	s = strings.ToLower(s)
	// Replace non-alphanumeric characters with hyphens
	reg := regexp.MustCompile("[^a-z0-9]+")
	s = reg.ReplaceAllString(s, "-")
	// Trim hyphens from start and end
	s = strings.Trim(s, "-")
	return s
}
