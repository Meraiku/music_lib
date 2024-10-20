package info

type InfoRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type InfoResponse struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}
