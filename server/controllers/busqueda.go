package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/Eddy150893/apiTestGo/server/services"
	"github.com/Eddy150893/apiTestGo/server/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	criterio := params["criterio"]
	if(criterio!=""){
		itunes:=services.GetItunes(services.Itunes,services.ItunesService,criterio)
		tvmaze:=services.GetTvMaze(services.Tvmaze,services.TvmazeService,criterio)
		crcind:=services.GetSoapDemo(services.Crcind,services.CrcindService,criterio)
		var response=models.Response{itunes,tvmaze,crcind,}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
	
}
