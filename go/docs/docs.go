// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/matrix/factorize": {
            "post": {
                "description": "Recibe una matriz y retorna dos matrices Q y R",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Factorial QR de una matriz",
                "operationId": "factorize-matrix",
                "parameters": [
                    {
                        "description": "Matriz a factorizar",
                        "name": "matrix",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "type": "number"
                                }
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.QRResponse"
                        }
                    }
                }
            }
        },
        "/matrix/rotate": {
            "post": {
                "description": "Recibe una matriz y una cantidad de grados de giro (múltiplos de 90) y retorna la matriz rotada dicha cantidad de grados",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Rota una matriz",
                "operationId": "rotate-matrix",
                "parameters": [
                    {
                        "description": "Matriz a rotar",
                        "name": "matrix",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "type": "number"
                                }
                            }
                        }
                    },
                    {
                        "description": "Grados de rotación",
                        "name": "degrees",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RotateResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.QRResponse": {
            "type": "object",
            "properties": {
                "Q": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "number"
                        }
                    }
                },
                "R": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "models.RotateResponse": {
            "type": "object",
            "properties": {
                "degrees": {
                    "type": "integer"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "type": "number"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/matrix",
	Schemes:          []string{},
	Title:            "Go API",
	Description:      "API de operaciones sobre matrices.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}