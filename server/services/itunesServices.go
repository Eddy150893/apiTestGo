package services

import (
    "fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/Eddy150893/apiTestGo/server/models"
)

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