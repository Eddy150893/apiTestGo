package main

import (
	"log"
	"net/http"
	"github.com/Eddy150893/apiTestGo/routes"
)

func main() {
	router := routes.NewRouter()
	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
