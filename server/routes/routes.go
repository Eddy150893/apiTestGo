package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Eddy150893/apiTestGo/server/controllers"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}
type Routes []Route

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

var routes = Routes{
	Route{
		"index",
		"GET",
		"/busqueda",
		controllers.Index,
	},
}
