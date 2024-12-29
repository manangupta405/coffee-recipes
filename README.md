# Coffee Recipes API

This project, **Coffee Recipes API**, is a learning-focused implementation of a web service that provides coffee recipes and suggestions. The project demonstrates the use of the Go programming language, modern web frameworks, and integration with AI services for generating content dynamically.

## Project Structure

```
coffee-recipes
├── Dockerfile           # Docker configuration for building and running the project
├── LICENSE              # Project license information
├── README.md            # Project documentation (this file)
├── cmd
│   └── server
│       └── main.go      # Main entry point for the application
├── config
│   ├── config.override.yaml # Override configuration for local customization
│   └── config.yaml          # Default configuration for the application
├── go.mod               # Go module dependencies
├── go.sum               # Dependency checksums
├── internal
│   ├── controllers
│   │   ├── receipe_controller.go       # Controller interface for recipes
│   │   └── receipe_controller_impl.go  # Controller implementation for recipes
│   ├── models
│   │   └── receipe.go                  # Model definition for coffee recipes
│   ├── repository
│   │   └── receipe_repository.go       # Repository interface for database interactions (to be implemented)
│   └── routes
│       └── routes.go                   # Route definitions for the API
└── pkg
    ├── ai
    │   ├── ai.go                       # AI client interface
    │   └── open_ai_impl.go             # OpenAI client implementation
    └── utils
        └── response.go                 # Utility functions for API responses
```

## Features

- **AI-Driven Coffee Recipes**: Uses OpenAI GPT models to generate coffee recipes and suggestions based on user-provided ingredients.
- **RESTful API**: Built using [Gin](https://gin-gonic.com/) for handling HTTP requests.
- **Modular Design**: Organized into controllers, models, routes, and utilities for scalability and maintainability.
- **Dockerized**: Easy to build and run using Docker.

## Setup and Usage

### Prerequisites

- [Go](https://golang.org/) 1.23 or higher
- [Docker](https://www.docker.com/) (optional but recommended)

### Running the Application

#### With Docker

1. Build the Docker image:
   ```sh
   docker build -t coffee-recipes .
   ```
2. Run the container:
   ```sh
   docker run -p 8080:8080 coffee-recipes
   ```
3. The API will be accessible at `http://localhost:8080/api`.

#### Without Docker

1. Install dependencies:
   ```sh
   go mod download
   ```
2. Run the server:
   ```sh
   go run cmd/server/main.go
   ```
3. The API will be accessible at `http://localhost:8080/api`.

### API Endpoints

- `POST /api/getPossibleCoffee`: Suggests coffee styles based on ingredients.
- `POST /api/getRecipe`: Provides a detailed recipe for a specific coffee style.
- `GET /api/health`: Health check endpoint.

## Disclaimer

This project is created **for learning purposes only**. The code has been partially or fully generated using OpenAI's ChatGPT. The author does not claim ownership of the code and assumes no liability for its use. Please do not use this project in production environments.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- [OpenAI](https://openai.com/) for the AI-generated content.
- The Go and Gin communities for excellent tools and frameworks.
