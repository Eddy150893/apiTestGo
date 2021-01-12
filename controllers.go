package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request) {
	query:=r.URL.Query()
	criterio:=query.Get("criterio")
	fmt.Println(criterio)
	if(criterio!=""){
		itunes:=GetItunes(Itunes,ItunesService,criterio)
		tvmaze:=GetTvMaze(Tvmaze,TvmazeService,criterio)
		crcind:=GetSoapDemo(Crcind,CrcindService,criterio)
		var response=Response{itunes,tvmaze,crcind,}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}
	
}
