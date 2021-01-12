package models
type ResultTvMaze struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

// TvmazeResponse struct ...
type DataTvmaze struct {
	Score float32 `json:"score"`
	Show  ResultTvMaze  `json:"show"`
}
