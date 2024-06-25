import { calculateMatrixStats } from "../services/matrix-stats.js";

// Obtener todos los usuarios
const analyzeMatrixHandler = (req, res) => {
  const { matrix } = req.body;

  if (!matrix) {
    return res.status(400).json({ error: "Campo 'matrix' requerido" });
  }

  try {
    res.json(calculateMatrixStats(matrix));
  } catch (error) {
    res.status(500).json({ error: "Error al calcular estadísticas de las matrices" });
  }
};

export { analyzeMatrixHandler };
