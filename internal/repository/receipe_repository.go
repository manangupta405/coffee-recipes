// internal/repository/recipe_repository.go
package repository

import "coffee-recipes/internal/models"

// RecipeRepository defines the interface for interacting with the database for recipe-related operations.
type RecipeRepository interface {
	// FetchAll retrieves all coffee recipes from the database.
	FetchAll() ([]*models.CoffeeRecipe, error)

	// FetchByID retrieves a specific coffee recipe by its ID.
	FetchByID(id string) (*models.CoffeeRecipe, error)

	// Create adds a new coffee recipe to the database.
	Create(recipe *models.CoffeeRecipe) error

	// Update modifies an existing coffee recipe in the database.
	Update(recipe *models.CoffeeRecipe) error

	// Delete removes a coffee recipe from the database by its ID.
	Delete(id string) error
}

// Note: Implement this interface with a concrete struct if database operations are required.
