import { Router } from "express";
import { analyzeMatrixHandler } from "../handlers/stats.js";

const router = Router();

// Ruta para obtener todos los usuarios
router.post("/stats", analyzeMatrixHandler);

export default router;
