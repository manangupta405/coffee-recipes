// internal/models/recipe.go
package models

// CoffeeRecipe represents the structure of a coffee recipe.
type CoffeeRecipe struct {
	// ID is the unique identifier for the coffee recipe.
	ID string `json:"id"`

	// Name is the name of the coffee recipe.
	Name string `json:"name" binding:"required"`

	// Ingredients is a list of ingredients required for the coffee recipe.
	// It must contain at least one ingredient.
	Ingredients []string `json:"ingredients" binding:"required,min=1"`

	// Instructions provides the steps to prepare the coffee.
	Instructions string `json:"instructions" binding:"required"`

	// Price is the cost associated with making the coffee recipe.
	// It must be greater than zero.
	Price float64 `json:"price" binding:"required,gt=0"`
}
