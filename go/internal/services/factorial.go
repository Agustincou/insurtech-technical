package services

import (
	"errors"
	"log"

	"gonum.org/v1/gonum/mat"
)

// Función para convertir una matriz de slice para JSON a una matriz Gonum
func SliceToMat(data [][]float64) (*mat.Dense, error) {
	// Validar que todas las filas tengan la misma longitud
	if !IsValidMatrix(data) {
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
func IsValidMatrix(data [][]float64) bool {
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

// O(n^2) -> Complejidad computacional por for anidados. Mejorable
// RotateMatrix rota una matriz MxN según el ángulo especificado.
func RotateMatrix(matrix [][]float64, angle int) ([][]float64, int, error) {
	m := len(matrix)    // Número de filas
	n := len(matrix[0]) // Número de columnas

	// Normalizar el ángulo a un rango de 0 a 360
	angle = ((angle % 360) + 360) % 360

	switch angle {
	case 0:
		return matrix, angle, nil

	case 90:
		rotated := make([][]float64, n)
		for i := range rotated {
			rotated[i] = make([]float64, m)
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				rotated[j][m-i-1] = matrix[i][j]
			}
		}

		return rotated, angle, nil

	case 180:
		rotated := make([][]float64, m)
		for i := range rotated {
			rotated[i] = make([]float64, n)
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				rotated[m-i-1][n-j-1] = matrix[i][j]
			}
		}

		return rotated, angle, nil

	case 270:
		rotated := make([][]float64, n)
		for i := range rotated {
			rotated[i] = make([]float64, m)
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				rotated[n-j-1][i] = matrix[i][j]
			}
		}

		return rotated, angle, nil

	default:
		log.Println("Ángulo no soportado")

		return nil, angle, errors.New("unsupported rotation degrees")
	}
}
