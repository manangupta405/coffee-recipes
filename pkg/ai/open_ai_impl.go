package ai

import (
	"coffee-recipes/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type openaiClient struct {
	oaiClient *openai.Client
}

func (o *openaiClient) GetPossibleCoffees(ctx context.Context, ingredients []string) ([]string, error) {
	// Step 1: Convert the ingredients list into JSON format
	ingredientBytes, err := json.Marshal(ingredients)
	if err != nil {
		return nil, fmt.Errorf("failed to convert ingredients to JSON: %w", err)
	}

	// Step 2: Create the prompt that instructs the AI on how to process the input
	prompt := `You are tasked with analyzing the provided ingredients and determining which coffee styles can be made.
Here are the instructions:
1. Identify coffee styles (e.g., Espresso, Latte, Mocha) that can be prepared using the ingredients provided.
2. Do not include the raw ingredient names as coffee styles.
3. Format your response strictly in the following JSON format:
{
  "coffees": ["Espresso", "Latte", "Cappuccino"],
  "failed": false
}`

	// Step 3: Define the expected format of the response using JSON schema
	coffeeParams := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"coffees": {
				Type:        jsonschema.Array,
				Description: "A list of coffee styles that can be prepared based on the given ingredients.",
				Items: &jsonschema.Definition{
					Type:        jsonschema.String,
					Description: "The name of a single coffee style, such as Espresso or Latte.",
				},
			},
			"failed": {
				Type:        jsonschema.Boolean,
				Description: "Indicates if the task could not be completed successfully.",
			},
		},
		Required: []string{"coffees", "failed"},
	}

	// Step 4: Register a function and tool definition for OpenAI to understand the task
	function := openai.FunctionDefinition{
		Name:        "get_possible_coffee_styles",
		Description: "Analyze the ingredients and generate a list of possible coffee styles.",
		Parameters:  coffeeParams,
	}
	tool := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &function,
	}

	// Step 5: Create the messages that define the system behavior and user input
	dialogue := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: prompt, // This message provides clear instructions to the AI on processing the ingredients
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: string(ingredientBytes), // The user's input is the list of ingredients in JSON format
		},
	}

	// Step 6: Call the OpenAI API to process the input and generate a response
	resp, err := o.oaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       "gpt-4",
		Messages:    dialogue,
		Tools:       []openai.Tool{tool},
		Temperature: 0, // Temperature set to 0 for consistent and deterministic results
		ToolChoice: &openai.ToolChoice{
			Type: openai.ToolTypeFunction,
			Function: openai.ToolFunction{
				Name: "get_possible_coffee_styles",
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to call OpenAI API: %w", err)
	}

	// Step 7: Extract the tool call results from the response
	if len(resp.Choices) == 0 || len(resp.Choices[0].Message.ToolCalls) == 0 {
		return nil, fmt.Errorf("no tool calls found in response")
	}

	firstToolCall := resp.Choices[0].Message.ToolCalls[0]
	if firstToolCall.Function.Name != "get_possible_coffee_styles" {
		return nil, fmt.Errorf("unexpected function called: %s", firstToolCall.Function.Name)
	}

	// Step 8: Parse the JSON response from the tool call into a structured format
	var result struct {
		Coffees []string `json:"coffees"`
		Failed  bool     `json:"failed"`
	}
	if err := json.Unmarshal([]byte(firstToolCall.Function.Arguments), &result); err != nil {
		return nil, fmt.Errorf("failed to parse tool call arguments: %w", err)
	}

	// Step 9: Handle cases where the AI could not generate valid coffee styles
	if result.Failed {
		return nil, fmt.Errorf("unable to generate coffee styles with the provided ingredients")
	}

	// Step 10: Filter out results that are direct matches to the input ingredients
	validCoffees := []string{}
	for _, coffee := range result.Coffees {
		isIngredient := false
		for _, ingredient := range ingredients {
			if strings.EqualFold(coffee, ingredient) {
				isIngredient = true
				break
			}
		}
		if !isIngredient {
			validCoffees = append(validCoffees, coffee)
		}
	}

	// Step 11: Ensure there are valid coffee styles to return, otherwise report an error
	if len(validCoffees) == 0 {
		return nil, fmt.Errorf("AI returned only ingredients or invalid results")
	}

	// Step 12: Return the final list of valid coffee styles
	return validCoffees, nil
}

func (o *openaiClient) GetRecipe(ctx context.Context, coffeeStyle string) (models.CoffeeRecipe, error) {
	// Step 1: Convert the coffee style into JSON format for the AI input
	styleBytes, err := json.Marshal(coffeeStyle)
	if err != nil {
		return models.CoffeeRecipe{}, fmt.Errorf("failed to convert coffee style to JSON: %w", err)
	}

	// Step 2: Create a prompt that guides the AI to generate a detailed recipe
	prompt := `You are tasked with creating a complete recipe for the specified coffee style.
Here are the instructions:
1. Assign a unique identifier to the recipe.
2. List all necessary ingredients with clear names.
3. Provide a detailed step-by-step guide for preparation.
4. Suggest a reasonable price for the coffee.
Respond strictly in the following JSON format:
{
  "id": "unique-id",
  "name": "Latte",
  "ingredients": ["Milk", "Espresso"],
  "instructions": "Step 1: Heat the milk. Step 2: Brew espresso. Step 3: Combine and serve.",
  "price": 4.50,
  "failed": false
}`

	// Step 3: Define the expected format of the recipe response using JSON schema
	recipeParams := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"id": {
				Type:        jsonschema.String,
				Description: "A unique identifier for the recipe.",
			},
			"name": {
				Type:        jsonschema.String,
				Description: "The name of the coffee style.",
			},
			"ingredients": {
				Type:        jsonschema.Array,
				Description: "A list of ingredients required to make the coffee.",
				Items: &jsonschema.Definition{
					Type:        jsonschema.String,
					Description: "The name of a single ingredient.",
				},
			},
			"instructions": {
				Type:        jsonschema.String,
				Description: "Step-by-step instructions for preparing the coffee.",
			},
			"price": {
				Type:        jsonschema.Number,
				Description: "The suggested price for the coffee.",
			},
			"failed": {
				Type:        jsonschema.Boolean,
				Description: "Indicates if the task could not be completed successfully.",
			},
		},
		Required: []string{"id", "name", "ingredients", "instructions", "price", "failed"},
	}

	// Step 4: Register a function and tool definition for OpenAI to understand the task
	function := openai.FunctionDefinition{
		Name:        "get_coffee_recipe",
		Description: "Generate a recipe for a specified coffee style.",
		Parameters:  recipeParams,
	}
	tool := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: &function,
	}

	// Step 5: Create the messages that define the system behavior and user input
	dialogue := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: prompt, // This message provides clear instructions to the AI for generating the recipe
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: string(styleBytes), // The user's input is the coffee style in JSON format
		},
	}

	// Step 6: Call the OpenAI API to process the input and generate a response
	resp, err := o.oaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       "gpt-4",
		Messages:    dialogue,
		Tools:       []openai.Tool{tool},
		Temperature: 0, // Temperature set to 0 for consistent and deterministic results
		ToolChoice: &openai.ToolChoice{
			Type: openai.ToolTypeFunction,
			Function: openai.ToolFunction{
				Name: "get_coffee_recipe",
			},
		},
	})
	if err != nil {
		return models.CoffeeRecipe{}, fmt.Errorf("failed to call OpenAI API: %w", err)
	}

	// Step 7: Extract the tool call results from the response
	if len(resp.Choices) == 0 || len(resp.Choices[0].Message.ToolCalls) == 0 {
		return models.CoffeeRecipe{}, fmt.Errorf("no tool calls found in response")
	}

	firstToolCall := resp.Choices[0].Message.ToolCalls[0]
	if firstToolCall.Function.Name != "get_coffee_recipe" {
		return models.CoffeeRecipe{}, fmt.Errorf("unexpected function called: %s", firstToolCall.Function.Name)
	}

	// Step 8: Parse the JSON response from the tool call into a structured format
	var result struct {
		ID           string   `json:"id"`
		Name         string   `json:"name"`
		Ingredients  []string `json:"ingredients"`
		Instructions string   `json:"instructions"`
		Price        float64  `json:"price"`
		Failed       bool     `json:"failed"`
	}
	if err := json.Unmarshal([]byte(firstToolCall.Function.Arguments), &result); err != nil {
		return models.CoffeeRecipe{}, fmt.Errorf("failed to parse tool call arguments: %w", err)
	}

	// Step 9: Handle cases where the AI could not generate a valid recipe
	if result.Failed {
		return models.CoffeeRecipe{}, fmt.Errorf("unable to generate recipe for the provided coffee style")
	}

	// Step 10: Validate the completeness and validity of the recipe data
	if len(result.Ingredients) == 0 || result.Instructions == "" || result.Price <= 0 {
		return models.CoffeeRecipe{}, fmt.Errorf("AI returned incomplete or invalid recipe data")
	}

	// Step 11: Return the final recipe details as a structured object
	return models.CoffeeRecipe{
		ID:           result.ID,
		Name:         result.Name,
		Ingredients:  result.Ingredients,
		Instructions: result.Instructions,
		Price:        result.Price,
	}, nil
}

func NewOpenAIClient(oaiClient *openai.Client) AIClient {
	// Create a new instance of the OpenAI client
	return &openaiClient{
		oaiClient: oaiClient,
	}
}
