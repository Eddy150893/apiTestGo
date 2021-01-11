package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/gorilla/mux"
	//"github.com/Eddy150893/apiTestGo/server/response"
	"github.com/Eddy150893/apiTestGo/server/services"
	//"github.com/Eddy150893/apiTestGo/server/parser"
)
type Result struct {
    Type string `json:"type"`
}
func Index(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	criterio := params["criterio"]
	fmt.Println(criterio)
	if(criterio!=""){
		responseData:=services.GetItunes(services.Itunes,criterio)
		buf := new(bytes.Buffer)
		buf.ReadFrom(responseData.Body)
		fmt.Fprintln(w, buf) // successfully prints the buffer to w and confirms successful JSON request from remote server
		var MyResult []Result
		json.Unmarshal(buf.Bytes(), &MyResult)
		for l := range MyResult {
			fmt.Fprintln(w, MyResult[l].Type)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(MyResult)
	}
	
}