package main

import (
	"log"

	"github.com/Agustincou/tec-exam/internal/routes"
)

func main() {
	r := routes.SetupRouter()

	// Iniciar el servidor en el puerto 8080
	log.Fatal(r.Listen(":8080"))
}
