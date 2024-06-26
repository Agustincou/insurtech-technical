package routes

import (
	"github.com/Agustincou/tec-exam/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, basePath string) {
	matrixApi := app.Group(basePath)

	// Rutas
	matrixApi.Post("/factorize", handlers.Factorize)
	matrixApi.Post("/rotate", handlers.Rotate)
}
