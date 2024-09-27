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

func (req *ModifySongRequest) Validate() map[string]string {
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
