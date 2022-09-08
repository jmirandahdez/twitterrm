package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/jmirandahdez/twitterrm/middlew"
	"github.com/jmirandahdez/twitterrm/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores seteo de puerto, el handler y poner en eschucha el Server
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("registro", middlew.ChekDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
