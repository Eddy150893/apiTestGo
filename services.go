package services

import (
    "fmt"
	"io/ioutil"
	"net/url"
	"net/http"
	"encoding/xml"
	"encoding/json"
	"github.com/Eddy150893/apiTestGo/models"
)

//Soap Services
var Crcind string="http://www.crcind.com"
var CrcindService string="/csp/samples/SOAP.Demo.cls/" 

//Rest Services
var Itunes string="https://itunes.apple.com/"
var ItunesService string="search?term="

var Tvmaze string="http://api.tvmaze.com/" 
var TvmazeService string="search/shows?q="

//Consulta a API crcind
func GetSoapDemo(apiUrl,servicio,criterio string)models.CrcResponse{
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
		fmt.Println("PETICION FALLIDA")
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var crcResponse models.CrcResponse
	xml.Unmarshal(body, &crcResponse)
	return crcResponse
}


//Consulta a API tvmaze
func GetTvMaze(apiUrl,servicio,criterio string) []models.DataTvmaze{
	response, err := http.Get(apiUrl+servicio+criterio)
	if err!=nil{
		fmt.Print(err)
		panic(err)
	}
	defer response.Body.Close()
	responseData,err:=ioutil.ReadAll(response.Body)

	var dataResult []models.DataTvmaze
	errMarshal:=json.Unmarshal(responseData,&dataResult)
	if errMarshal !=nil{
		fmt.Print("Error en marshal")
		panic(errMarshal)
	}
	return dataResult
}

//Consulta a API itunes
func GetItunes(apiUrl,servicio,criterio string) models.DataItunes{
	response, err := http.Get(apiUrl+servicio+criterio)
	if err!=nil{
		fmt.Print(err.Error())
		panic(err)
	}
	defer response.Body.Close()
	responseData,err:=ioutil.ReadAll(response.Body)
	var dataResult models.DataItunes
	errMarshal:=json.Unmarshal(responseData,&dataResult)
	if errMarshal !=nil{
		fmt.Print(errMarshal.Error())
		panic(errMarshal)
	}
	return dataResult
}