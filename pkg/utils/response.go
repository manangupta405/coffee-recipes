// pkg/utils/response.go
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONResponse sends a JSON response with a given status code and payload.
func JSONResponse(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

// SuccessResponse sends a success response with HTTP 200 status.
// Parameters:
// - c: The Gin context.
// - data: The data to include in the response payload.
func SuccessResponse(c *gin.Context, data interface{}) {
	JSONResponse(c, http.StatusOK, data)
}

// ErrorResponse sends an error response with a custom HTTP status code.
// Parameters:
// - c: The Gin context.
// - status: The HTTP status code for the error.
// - message: The error message to include in the response.
func ErrorResponse(c *gin.Context, status int, message string) {
	JSONResponse(c, status, gin.H{"error": message})
}
