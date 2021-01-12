package models

type Response struct{
	Itunes DataItunes `json:"itunes"`
	Tvmaze []DataTvmaze `json:"tvmaze"`
	Crcind CrcResponse `json:"crcind"`
}