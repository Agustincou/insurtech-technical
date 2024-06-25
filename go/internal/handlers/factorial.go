package handlers

import (
	"log"

	"github.com/Agustincou/tec-exam/internal/errors"
	"github.com/Agustincou/tec-exam/internal/services"

	"github.com/gofiber/fiber/v2"
	"gonum.org/v1/gonum/mat"
)

//ToDo: Logueo estructurado con request-id y tracing de línea de código

func Factorize(c *fiber.Ctx) error {
	// Decodificar la matriz recibida en formato JSON
	var inputMatrix struct {
		Value [][]float64 `json:"matrix"`
	}

	if err := c.BodyParser(&inputMatrix); err != nil {
		log.Println(err.Error())

		return errors.JSONRequestEnconding
	}

	// Convertir la matriz de entrada a una matriz Gonum
	matrixToProcess, matrixErr := services.SliceToMat(inputMatrix.Value)

	if matrixErr != nil {
		log.Println(matrixErr.Error())

		return errors.InvalidMatrix
	}

	// Calcular la factorización QR
	var qr mat.QR
	qr.Factorize(matrixToProcess)

	// Obtener matrices Q y R
	var Q mat.Dense
	var R mat.Dense
	qr.QTo(&Q)
	qr.RTo(&R)

	// Convertir matrices Q y R a estructuras de slice para JSON
	qrResult := struct {
		Q [][]float64 `json:"Q"`
		R [][]float64 `json:"R"`
	}{
		Q: services.MatToSlice(&Q),
		R: services.MatToSlice(&R),
	}

	if err := c.BodyParser(&qrResult); err != nil {
		return errors.JSONResponseEnconding
	}

	// Codificar qrResult a JSON y enviar como respuesta
	return c.Status(fiber.StatusOK).JSON(qrResult)
}
