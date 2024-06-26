package main

import (
	"log"

	_ "github.com/Agustincou/tec-exam/docs"
	"github.com/Agustincou/tec-exam/internal/errors"
	"github.com/Agustincou/tec-exam/internal/routes"

	"github.com/gofiber/fiber/v2"

	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "github.com/swaggo/files"
)

// @title Go API
// @version 1.0
// @description API de operaciones sobre matrices.
// @host localhost:8080
// @BasePath /matrix
func main() {
	app := fiber.New()

	// Ruta para la documentación de Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

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

	routes.Setup(app, "/matrix")

	log.Fatal(app.Listen(":8080"))
}
