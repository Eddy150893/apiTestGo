package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/Eddy150893/apiTestGo/services"
	"github.com/Eddy150893/apiTestGo/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	query:=r.URL.Query()
	criterio:=query.Get("criterio")
	fmt.Println(criterio)
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
