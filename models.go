package main

import (
	"encoding/xml"
)
//Response estructura para respuesta estandar de la API
type Response struct{
	Itunes []ResultItunes `json:"itunes"`
	Tvmaze []DataTvmaze `json:"tvmaze"`
	Crcind []PersonIdentification `json:"crcind"`
}


//ResultItunes estructura con los campos a obtener de la API Itunes
//DataItunes estructura utilizada para mapear los datos provenientes de itunes
type ResultItunes struct{
	Kind string `json:"kind"`
	WrapperType string `json:"WrapperType"`
	ArtisName string `json:"artistName"`
	TrackName string `json:"trackName"`
	CollectionName string `json:"collectionName"`
	PreviewURL string `json:"previewUrl"`
	TrackViewURL string `json:"trackViewUrl"`
}
type DataItunes struct{
	ResultCount int `json:"resultCount"`
	Results []ResultItunes
}

//ResultTvMaze estructura con los campos a obtener de la API tvmaze
//DataTvmaze estructura utilizada para mapear los datos provenientes de tvmaze
type Country struct{
	Name string `json:"name"`
	Code string `json:"code"`
	Timezone string `json:"timezone"`
}

type Network struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Country Country `json:"country"`
}

type Image struct{
	Medium string `json:"medium"`
	Original string `json:"original"`
}

type ResultTvMaze struct {
	Id  int `json:"id"`
	Name string `json:"name"`
	Image Image `json:"image"`
	Network Network `json:"network"`
	Language string `json:"language"`
}
type DataTvmaze struct {
	Score float32 `json:"score"`
	Show  ResultTvMaze  `json:"show"`
}


//PersonIdentification estructura con los campos a obtener de la API crcind
type PersonIdentification struct {
	XMLName xml.Name `xml:"PersonIdentification"`
	ID      int      `xml:"ID"`
	Name    string   `xml:"Name"`
	SSN     string   `xml:"SSN"`
	DOB     string   `xml:"DOB"`
}

//GetListByNameResult estructura utilizada en GetListByNameResponse struct
type GetListByNameResult struct {
	XMLName              xml.Name               `xml:"GetListByNameResult"`
	PersonIdentification []PersonIdentification `xml:"PersonIdentification"`
}

//GetListByNameResponse estructura utilizada en el Body struct 
type GetListByNameResponse struct {
	XMLName             xml.Name            `xml:"GetListByNameResponse"`
	GetListByNameResult GetListByNameResult `xml:"GetListByNameResult"`
}

//Body estructura utilizada para mapear el body del CrcResponse 
type Body struct {
	XMLName               xml.Name
	GetListByNameResponse GetListByNameResponse `xml:"GetListByNameResponse"`
}


type CrcResponse struct {
	XMLName xml.Name
	Body    Body
}




