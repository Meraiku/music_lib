package response

type GetSongsResponse struct {
	ID    string `json:"id"`
	Group string `json:"group"`
	Song  string `json:"song"`
}
