// internal/controllers/recipe_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
)

// RecipeController defines the interface for handling recipe-related operations.
type RecipeController interface {
	// GetPossibleCoffee processes a request to fetch possible coffee combinations based on ingredients.
	GetPossibleCoffee(c *gin.Context)

	// GetRecipe processes a request to fetch the recipe for a specific coffee type.
	GetRecipe(c *gin.Context)

	// HealthCheck verifies the health and readiness of the service.
	HealthCheck(c *gin.Context)
}
