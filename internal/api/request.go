package api

type AddSongRequest struct {
	Group string `json:"group" example:"Muse"`
	Song  string `json:"song" example:"Supermassive Black Hole"`
}

func (req *AddSongRequest) Validate() map[string]string {
	errors := map[string]string{}

	if req.Group == "" {
		errors["group"] = "is not provided"
	}
	if req.Song == "" {
		errors["song"] = "is not provided"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

type PatchRequest struct {
	Group       *string `json:"group" example:"Rick Astley"`
	Song        *string `json:"song" example:"Never Gonna Give You Up"`
	ReleaseDate *string `json:"releaseDate" example:"24.07.1987"`
	Text        *string `json:"text" example:"We’re no strangers to love\nYou know the rules and so do I\nA full commitment’s what I’m thinking of\nYou wouldn’t get this from any other guy\n\nI just wanna tell you how I’m feeling\nGotta make you understand\n\nNever gonna give you up\nNever gonna let you down\nNever gonna run around and desert you\nNever gonna make you cry\nNever gonna say goodbye\nNever gonna tell a lie and hurt you\n\nWe’ve known each other for so long\nYour heart’s been aching but you’re too shy to say it\nInside we both know what’s been going on\nWe know the game and we’re gonna play it\n\nAnd if you ask me how I’m feeling\nDon’t tell me you’re too blind to see\n\nNever gonna give you up\nNever gonna let you down\nNever gonna run around and desert you\nNever gonna make you cry\nNever gonna say goodbye\nNever gonna tell a lie and hurt you"`
	Link        *string `json:"link" example:"https://www.youtube.com/watch?v=dQw4w9WgXcQ"`
}
