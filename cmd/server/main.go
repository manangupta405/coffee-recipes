// cmd/server/main.go
package main

import (
	"coffee-recipes/internal/controllers"
	"coffee-recipes/internal/routes"
	"coffee-recipes/pkg/ai"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func main() {
	// Initialize configuration
	if err := initConfig(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	} else {
		log.Println("Configuration loaded successfully")
	}

	// Set Gin to release mode if specified in the configuration
	if viper.GetString("mode") == "release" {
		gin.SetMode(gin.ReleaseMode)
		log.Println("Running in release mode")
	} else {
		log.Println("Running in debug mode")
	}

	// Initialize AI Client
	apiKey := viper.GetString("openai.api_key")
	if apiKey == "" {
		log.Fatal("OpenAI API key is not configured")
	}
	log.Println("Initializing OpenAI client")
	aiClient := ai.NewOpenAIClient(openai.NewClient(apiKey))

	// Create Recipe Controller
	log.Println("Creating Recipe Controller")
	recipeController := controllers.NewRecipeController(aiClient)

	// Initialize Gin router
	log.Println("Initializing Gin router")
	router := gin.Default()

	// Setup routes with the RecipeController
	log.Println("Setting up routes")
	routes.SetupRoutes(router, recipeController)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("server.port")
		if port == "" {
			log.Fatal("Server port is not configured")
		}
	}
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func initConfig() error {
	// Load the default configuration from config.yaml
	log.Println("Loading default configuration from config.yaml")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	log.Println("Default configuration loaded successfully")

	// Check if config.override.yaml exists and load it
	overrideConfigPath := "./config/config.override.yaml"
	if _, err := os.Stat(overrideConfigPath); err == nil {
		log.Printf("Override config file found: %s", overrideConfigPath)
		viper.SetConfigFile(overrideConfigPath)
		if err := viper.MergeInConfig(); err != nil {
			log.Printf("Warning: Unable to merge override config: %v", err)
		} else {
			log.Println("Override config loaded successfully")
		}
	} else {
		log.Println("No override config file found")
	}

	return nil
}
