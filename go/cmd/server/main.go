package main

import (
	"log"

	"github.com/Agustincou/tec-exam/internal/routes"
)

func main() {
	r := routes.SetupRouter()

	log.Fatal(r.Listen(":8080"))
}
