package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/michee/pkg/routes"
)

const port = ":5000"

func main() {
	r := mux.NewRouter()
	// routes.RegisterHotel(r)
	routes.RegisterUser(r)
	http.Handle("/", r)

	fmt.Printf("le serveur fonctionne sur http://localhost%s", port)

	log.Fatal(http.ListenAndServe("localhost:5000", r))
}