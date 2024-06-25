import express from "express";
import bodyParser from "body-parser";
import routes from "./routes/routes.js";

const app = express();
const port = 3000;

// Middleware para parsear application/json
app.use(bodyParser.json());

// Rutas
app.use("/matrix", routes);

// Manejador de ruta para errores 404
app.use((req, res, next) => {
  res.status(404).json({ message: "Route not found" });
});

// Manejador de errores global
app.use((err, req, res, next) => {
  console.error(err.stack);
  res.status(500).json({ message: "Internal Server Error" });
});

// Iniciar el servidor
app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});
