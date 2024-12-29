// internal/routes/routes.go
package routes

import (
	"coffee-recipes/internal/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API endpoints and maps them to their respective controller methods.
func SetupRoutes(router *gin.Engine, recipeController controllers.RecipeController) {
	// Group API endpoints under /api
	api := router.Group("/api")
	{
		// Endpoint to get possible coffee combinations based on ingredients
		api.POST("/getPossibleCoffee", recipeController.GetPossibleCoffee)

		// Endpoint to get the recipe for a specific coffee type
		api.POST("/getRecipe", recipeController.GetRecipe)

		// Health Check Endpoint to verify service readiness
		api.GET("/health", recipeController.HealthCheck)
	}
}
