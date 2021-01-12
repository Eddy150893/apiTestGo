package models

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
