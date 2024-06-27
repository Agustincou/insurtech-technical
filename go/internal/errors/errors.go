package errors

import (
	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	HttpStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

func (a *APIError) Error() string {
	return a.Message
}

var (
	JSONRequestEnconding  = &APIError{HttpStatus: fiber.StatusBadRequest, Code: 1000, Message: "Error al codificar solicitud JSON"}
	JSONResponseEnconding = &APIError{HttpStatus: fiber.StatusInternalServerError, Code: 1001, Message: "Error al codificar respuesta JSON"}
	InvalidMatrix         = &APIError{HttpStatus: fiber.StatusBadRequest, Code: 1002, Message: "Matriz inválida"}
	InvalidRotation       = &APIError{HttpStatus: fiber.StatusBadRequest, Code: 1003, Message: "Rotación inválida"}
	InternalAPI           = &APIError{HttpStatus: fiber.StatusInternalServerError, Code: 1004, Message: "Internal API Error"}
)
