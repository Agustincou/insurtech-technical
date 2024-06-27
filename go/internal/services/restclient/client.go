package restclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Agustincou/tec-exam/internal/errors"
	"github.com/Agustincou/tec-exam/internal/models"
)

// Client estructura del cliente REST
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient crea una nueva instancia del cliente REST
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{},
	}
}

// PostMatrix envía una matriz al endpoint especificado
func (c *Client) PostMatrix(matrix [][]float64) (*models.MatrixStats, error) {
	// Datos a enviar en el body
	data := map[string]interface{}{
		"matrix": matrix,
	}

	// Convertir los datos a JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Crear el request
	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Establecer el header Content-Type a application/json
	req.Header.Set("Content-Type", "application/json")

	// Enviar el request usando el cliente HTTP
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.InternalAPI
	}

	// Decodificar la respuesta
	var stats models.MatrixStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		log.Println("Invalid stats API response format")

		return nil, err
	}

	return &stats, nil
}
