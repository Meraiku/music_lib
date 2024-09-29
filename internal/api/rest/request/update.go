package request

type PatchRequest struct {
	Group       *string `json:"group"`
	Song        *string `json:"song"`
	ReleaseDate *string `json:"releaseDate"`
	Text        *string `json:"text"`
	Link        *string `json:"link"`
}
