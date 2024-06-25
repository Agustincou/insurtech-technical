package routes

import (
	"log"

	"github.com/Agustincou/tec-exam/internal/errors"
	"github.com/Agustincou/tec-exam/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter() *fiber.App {
	app := fiber.New()

	// Definición del middleware para manejo de errores global
	app.Use(func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			log.Printf("Error: %v\n", err)

			if apiError, isAPIErr := err.(*errors.APIError); isAPIErr {
				return c.Status(apiError.HttpStatus).JSON(fiber.Map{
					"code":    apiError.Code,
					"message": apiError.Error(),
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return nil
	})

	matrixApi := app.Group("/matrix")

	// Rutas
	matrixApi.Post("/factorize", handlers.Factorize)
	matrixApi.Post("/rotate", handlers.Rotate)

	return app
}
