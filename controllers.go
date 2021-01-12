package main

import (
	"net/http"
	"encoding/json"
)

//Index() funcion utilizada en las ruta index de routes.go
//el pattern de la ruta es /busqueda?criterio=string

//Entrada :criterio de busqueda
//Salida :json con los resultados de busqueda
//en las apis de terceros

//Las busquedas se realizan a traves de diferentes funciones Get
//declaradas en services.go

//Response es una estructura declarada en models.go 
func Index(w http.ResponseWriter, r *http.Request) {
	//obtencion de criterio contenido en la url
	query:=r.URL.Query()
	criterio:=query.Get("criterio")

	//Validacion si el criterio viene vacio, No realiza busqueda
	if(criterio!=""){
		//Consumo de APIs de terceros
		itunes:=GetItunes(Itunes,ItunesService,criterio)
		tvmaze:=GetTvMaze(Tvmaze,TvmazeService,criterio)
		crcind:=GetSoapDemo(Crcind,CrcindService,criterio)

		//Unificacion de respuestas
		var response=Response{itunes.Results,tvmaze,crcind.Body.GetListByNameResponse.GetListByNameResult.PersonIdentification,}
		
		//Envio y codificacion(json) de la respuesta al cliente
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}else{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Debe ingresar un criterio de busqueda")
	}
	
}
