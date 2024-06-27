package models

// swagger:model
type Request struct {
	Value [][]float64 `json:"matrix"`

	// required: false
	WithStatistics bool `json:"with_statistics"`
}

type RotationRequest struct {
	Request
	RotationDegrees int `json:"degrees"`
}
