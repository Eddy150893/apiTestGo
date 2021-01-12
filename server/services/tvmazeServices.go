package services

import (
    "fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/Eddy150893/apiTestGo/server/models"
)

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