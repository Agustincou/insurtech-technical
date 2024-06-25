// Función para calcular estadísticas de una matriz
function calculateMatrixStats(matrix) {
  const flatMatrix = matrix.flat();

  const max = Math.max(...flatMatrix);
  const min = Math.min(...flatMatrix);
  const sum = flatMatrix.reduce((acc, curr) => acc + curr, 0);
  const avg = sum / flatMatrix.length;

  const is_diagonal = checkIfDiagonal(matrix);

  return {
    max,
    min,
    sum,
    avg,
    is_diagonal,
  };
}

// Función para verificar si una matriz es diagonal con margen de redondeo. Default 1e-10
function checkIfDiagonal(matrix, tolerance = 1e-10) {
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix[i].length; j++) {
      // Comparar usando un margen de redondeo
      if (i !== j && Math.abs(matrix[i][j]) > tolerance) {
        return false;
      }
    }
  }

  return true;
}

export { calculateMatrixStats, checkIfDiagonal };
