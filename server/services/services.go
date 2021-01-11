package services
import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

//Rest Services
var Itunes string="https://itunes.apple.com/search?term="
var Tvmaze string="http://api.tvmaze.com/search/shows?q=" 
//Soap Services
var Soapdemo string="http://www.crcind.com/csp/samples/SOAP.Demo.cls?wsdl" 

func GetItunes(servicio,criterio string)string{
	response, err := http.Get(servicio+criterio)
	if err!=nil{
		fmt.Print(err.Error())
		panic(err)

	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
        log.Fatal(err)
	}
	return string(responseData)
}