package services

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

// Función para convertir una matriz de slice para JSON a una matriz Gonum
func SliceToMat(data [][]float64) (*mat.Dense, error) {
	// Validar que todas las filas tengan la misma longitud
	if !validateMatrix(data) {
		return nil, errors.New("invalid matrix")
	}

	r := len(data)
	c := len(data[0])
	m := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m.Set(i, j, data[i][j])
		}
	}

	return m, nil
}

// Función para convertir una matriz Gonum a una matriz de slice para JSON
func MatToSlice(m *mat.Dense) [][]float64 {
	r, c := m.Dims()
	data := make([][]float64, r)
	for i := 0; i < r; i++ {
		data[i] = make([]float64, c)
		for j := 0; j < c; j++ {
			data[i][j] = m.At(i, j)
		}
	}

	return data
}

// Función auxiliar para validar si la matriz es válida
func validateMatrix(data [][]float64) bool {
	if len(data) == 0 {
		return false
	}

	cols := len(data[0])
	for i := 1; i < len(data); i++ {
		if len(data[i]) != cols {
			return false
		}
	}

	return true
}
