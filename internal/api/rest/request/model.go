package request

type AddSongRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type ModifySongRequest struct {
	ID    string `json:"id"`
	Group string `json:"group"`
	Song  string `json:"song"`
}
