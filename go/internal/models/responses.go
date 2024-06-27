package models

import "github.com/Agustincou/tec-exam/pkg/enum"

// QRResponse representa la respuesta para la factorización QR
type QRResponse struct {
	Results []MatrixResult `json:"results"`
}

// RotateResponse representa la respuesta para la rotación de una matriz
type RotateResponse struct {
	Degrees int            `json:"degrees"`
	Results []MatrixResult `json:"results"`
}

type MatrixResult struct {
	Type  enum.MatrixType `json:"type"`
	Value [][]float64     `json:"value"`
	Stats *MatrixStats    `json:"stats,omitempty"`
}

type MatrixStats struct {
	Max        float64 `json:"max"`
	Min        float64 `json:"min"`
	Sum        float64 `json:"sum"`
	Avg        float64 `json:"avg"`
	IsDiagonal bool    `json:"is_diagonal"`
}
