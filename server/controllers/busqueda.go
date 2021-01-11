package controllers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/Eddy150893/apiTestGo/server/response"
	"github.com/Eddy150893/apiTestGo/server/services"
)

func Index(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	criterio := params["criterio"]
	fmt.Println(criterio)
	if(criterio!=""){
		result:=services.GetItunes(services.Itunes,criterio)
		response.Response(w, 200, result)
	}
	
}