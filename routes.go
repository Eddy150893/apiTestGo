package main

import (
	"net/http"
	"github.com/gorilla/mux"
)
//Route estructura para definir rutas
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}
//Routes slices de las rutas del servidor
type Routes []Route

//NewRouter funcion de mapeo de rutas
//devuelve un tipo *mux.Router que contiene todas las rutas
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}

//Route index unica ruta dentro de la aplicacion
//Utiliza la accion Index del archivo controllers.go
var routes = Routes{
	Route{
		"index",
		"GET",
		"/busqueda",
		Index,
	},
}
