package models

// QRResponse representa la respuesta para la factorización QR
type QRResponse struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

// RotateResponse representa la respuesta para la rotación de una matriz
type RotateResponse struct {
	Degrees int         `json:"degrees"`
	Result  [][]float64 `json:"result"`
}
