package models

import (
	"encoding/xml"
)

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