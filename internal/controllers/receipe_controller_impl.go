// internal/controllers/recipe_controller_impl.go
package controllers

import (
	"coffee-recipes/pkg/ai"
	"coffee-recipes/pkg/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type recipeControllerImpl struct {
	aiClient ai.AIClient
}

// NewRecipeController creates a new instance of RecipeController with the provided AI client.
func NewRecipeController(aiClient ai.AIClient) RecipeController {
	return &recipeControllerImpl{aiClient: aiClient}
}

// GetPossibleCoffee handles the request to fetch possible coffee combinations based on provided ingredients.
func (r *recipeControllerImpl) GetPossibleCoffee(c *gin.Context) {
	// Parse the incoming request for ingredients
	var request struct {
		Ingredients []string `json:"ingredients" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Call AI client to get possible coffee combinations
	coffees, err := r.aiClient.GetPossibleCoffees(context.Background(), request.Ingredients)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the list of possible coffees
	utils.SuccessResponse(c, gin.H{"possible_coffees": coffees})
}

// GetRecipe handles the request to fetch a recipe for a specific coffee type.
func (r *recipeControllerImpl) GetRecipe(c *gin.Context) {
	// Parse the incoming request for coffee type
	var request struct {
		CoffeeType string `json:"coffee_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Call AI client to get the recipe for the specified coffee type
	recipe, err := r.aiClient.GetRecipe(context.Background(), request.CoffeeType)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the recipe details
	utils.SuccessResponse(c, gin.H{
		"recipe": gin.H{
			"ingredients":  recipe.Ingredients,
			"instructions": recipe.Instructions,
		},
	})
}

// HealthCheck handles the health check request.
func (r *recipeControllerImpl) HealthCheck(c *gin.Context) {
	// Respond with a simple readiness message
	utils.SuccessResponse(c, "Ready")
}
