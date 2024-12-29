package ai

import (
	"coffee-recipes/internal/models"
	"context"
)

// AIClient defines the interface for interacting with the AI system to retrieve coffee recipes and suggestions.
type AIClient interface {
	// GetPossibleCoffees fetches a list of possible coffee styles based on provided ingredients.
	// Parameters:
	// - ctx: Context for request control.
	// - ingredients: List of ingredients available for making coffee.
	// Returns:
	// - A list of coffee styles as strings.
	// - An error, if the operation fails.
	GetPossibleCoffees(ctx context.Context, ingredients []string) ([]string, error)

	// GetRecipe fetches the detailed recipe for a specific coffee style.
	// Parameters:
	// - ctx: Context for request control.
	// - coffeeStyle: Name of the coffee style for which the recipe is requested.
	// Returns:
	// - A CoffeeRecipe object containing ingredients and instructions.
	// - An error, if the operation fails.
	GetRecipe(ctx context.Context, coffeeStyle string) (models.CoffeeRecipe, error)
}
