package main

import (
	"log"
	"net/http"
	"os"
)
//main inicio de ejecucion y levantado del servidor
//NewRouter() funcion definida en routes.go
func main() {
	router := NewRouter()
	//port obtine el puerto del host de no encontrarlo
	//se setea el 8080
	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}
	server := http.ListenAndServe(":"+port, router)
	log.Fatal(server)
}
