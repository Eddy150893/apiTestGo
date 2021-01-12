package main

import (
    "fmt"
	"io/ioutil"
	"net/url"
	"net/http"
	"encoding/xml"
	"encoding/json"
)

//Itunes Url API itunes(rest)
//ItunesService Urn API itunes
var Itunes string="https://itunes.apple.com/"
var ItunesService string="search?term="

//Tvmaze Url API tvmaze(rest)
//TvmazeService Urn API tvmze
var Tvmaze string="http://api.tvmaze.com/" 
var TvmazeService string="search/shows?q="

//Crcind Url API crcind(soap)
//CrcindService Urn servicio API crcind
var Crcind string="http://www.crcind.com"
var CrcindService string="/csp/samples/SOAP.Demo.cls/" 


//GetItunes funcion que consume servicio de itunes por medio de GET
//Parametros Url de la API, Urn del servicio, un string de busqueda(criterio)
//Retorna un DataItunes con los datos obtenidos(models.go)
func GetItunes(apiUrl,servicio,criterio string) DataItunes{
	var dataResult DataItunes
	response, err := http.Get(apiUrl+servicio+criterio)
	if err!=nil{
		fmt.Print("Error en http.Get de Itunes")
		return dataResult
	}
	defer response.Body.Close()
	responseData,err:=ioutil.ReadAll(response.Body)
	
	errMarshal:=json.Unmarshal(responseData,&dataResult)
	if errMarshal !=nil{
		fmt.Print("Error en json.Unmarshal en itunes")
		return dataResult
	}
	return dataResult
}

//GetTvMaze funcion que consume servicio de tvmaze por medio de GET
//Parametros Url de la API, Urn del servicio, un string de busqueda(criterio)
//Retorna un []DataTvmaze con los datos obtenidos(models.go) 
func GetTvMaze(apiUrl,servicio,criterio string) []DataTvmaze{
	var dataResult []DataTvmaze
	response, err := http.Get(apiUrl+servicio+criterio)
	if err!=nil{
		fmt.Print("Error en http.Get de tvmaze")
		return dataResult
	}
	defer response.Body.Close()
	responseData,err:=ioutil.ReadAll(response.Body)
	errMarshal:=json.Unmarshal(responseData,&dataResult)
	if errMarshal !=nil{
		fmt.Print("Error en json.Unmarshal en tvmaze")
		return dataResult
	}
	return dataResult
}

//Consulta a API crcind
func GetSoapDemo(apiUrl,servicio,criterio string)CrcResponse{
	var crcResponse CrcResponse

	data := url.Values{}
	data.Set("name", criterio)
	data.Set("soap_method", "GetListByName")
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = servicio
	u.RawQuery = data.Encode()
	urlStr := u.String() 
	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, urlStr, nil) 
	r.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("error en client.Do")
		return crcResponse
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	errMarshal:=xml.Unmarshal(body, &crcResponse)
	if errMarshal !=nil{
		fmt.Print("Error en xml.Unmarshal en crcind")
		return crcResponse
	}
	return crcResponse
}




