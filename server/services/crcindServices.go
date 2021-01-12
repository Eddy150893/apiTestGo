package services
import (
    "fmt"
	"io/ioutil"
	"net/url"
	"net/http"
	"encoding/xml"
	"github.com/Eddy150893/apiTestGo/server/models"
)

func GetSoapDemo(apiUrl,servicio,criterio string)models.CrcResponse{
	data := url.Values{}
	data.Set("name", criterio)
	data.Set("soap_method", "GetListByName")
	// se parsea la url
	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = servicio
	// se codifican las variables en la query
	u.RawQuery = data.Encode()
	// se asigna la url completa a la variable final
	urlStr := u.String() // "https://itunes.apple.com/search/?limit=25&media=all&term=foo"
	// se crea el cliente que hara la peticion (se usa un puntero)
	client := &http.Client{}
	// se crea la request
	r, _ := http.NewRequest(http.MethodGet, urlStr, nil) // URL-encoded payload\
	// se asigna el header
	r.Header.Add("Content-Type", "application/json")

	// se ejecuta el request
	resp, err := client.Do(r)
	// handle error, la peticion fallo
	if err != nil {
		fmt.Println("PETICION FALLIDA")
		fmt.Println(err)
	}

	// despues del flujo cierra la respuesta
	defer resp.Body.Close()
	// leo todo el body
	body, _ := ioutil.ReadAll(resp.Body)

	var crcResponse models.CrcResponse
	// formateo el body de json a go y se asigna a appleResponse
	xml.Unmarshal(body, &crcResponse)

	return crcResponse
}