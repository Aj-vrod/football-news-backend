package models

type NewsV1Response struct {
	Results []NewsItem `json:"results"`
}

type NewsItem struct {
	URL      string `json:"url"`
	Title    string `json:"title"`
	Source   string `json:"source"`
	Creation string `json:"creation_date"`
}
