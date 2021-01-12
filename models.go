package models

import (
	"encoding/xml"
)

//========Modelo de Response
type Response struct{
	Itunes DataItunes `json:"itunes"`
	Tvmaze []DataTvmaze `json:"tvmaze"`
	Crcind CrcResponse `json:"crcind"`
}
//=========Fin de modelo Response

//=========Modelo de resultado API crcind
type PersonIdentification struct {
	XMLName xml.Name `xml:"PersonIdentification"`
	ID      int      `xml:"ID"`
	Name    string   `xml:"Name"`
	SSN     string   `xml:"SSN"`
	DOB     string   `xml:"DOB"`
}

// GetListByNameResult struct ...
type GetListByNameResult struct {
	XMLName              xml.Name               `xml:"GetListByNameResult"`
	PersonIdentification []PersonIdentification `xml:"PersonIdentification"`
}

// GetListByNameResponse struct ...
type GetListByNameResponse struct {
	XMLName             xml.Name            `xml:"GetListByNameResponse"`
	GetListByNameResult GetListByNameResult `xml:"GetListByNameResult"`
}

// Body struct ...
type Body struct {
	XMLName               xml.Name
	GetListByNameResponse GetListByNameResponse `xml:"GetListByNameResponse"`
}

// CrcResponse struct ...
type CrcResponse struct {
	XMLName xml.Name
	Body    Body
}

// StandardResponse struct ...
type StandardResponse struct {
	Category   string `json:"category"`
	Name       string `json:"name"`
	Author     string `json:"author"`
	PreviewURL string `json:"previewUrl"`
	Origin     string `json:"origin"`
}
//=========Fin modelo crcind

//=========Modelo de Resultado Api tvmaze
type ResultTvMaze struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// TvmazeResponse struct ...
type DataTvmaze struct {
	Score float32 `json:"score"`
	Show  ResultTvMaze  `json:"show"`
}
//=========Fin de modelo tvmaze

//=========Modelo de Resultado Api itunes
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
//=========Fin de modelo itunes