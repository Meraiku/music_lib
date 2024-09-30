package api

type Song struct {
	ID          string `json:"id" example:"uuid"`
	Group       string `json:"group" example:"Muse"`
	Song        string `json:"song" example:"Supermassive Black Hole"`
	ReleaseDate string `json:"releaseDate" example:"16.07.2006"`
	Text        string `json:"text" example:"oh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight"`
	Link        string `json:"link" example:"https://www.youtube.com/watch?v=Xsp3_a-PMTw"`
}

type Text struct {
	VerseNumber int    `json:"verse_number" example:"1"`
	Verse       string `json:"verse" example:"Never gonna give you up\n"`
}
