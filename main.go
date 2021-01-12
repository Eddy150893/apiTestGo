package main

import (
	"log"
	"net/http"
)
//main inicio de ejecucion y levantado del servidor
//NewRouter() funcion definida en routes.go
func main() {
	router := NewRouter()
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
