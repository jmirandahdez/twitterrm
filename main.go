package main

import (
	"log"

	"github.com/jmirandahdez/twitterrm/bd"
	"github.com/jmirandahdez/twitterrm/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
