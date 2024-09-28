package rest

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/meraiku/music_lib/internal/api/rest/request"
	"github.com/meraiku/music_lib/internal/converter"
	"github.com/meraiku/music_lib/internal/model"
	"github.com/meraiku/music_lib/internal/repo"
)

func statusCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("OK"))
}

// @Summary		Get Songs
// @Description	Prints List of songs
// @Tags			Songs
// @Produce		json
// @Success		200	{array}		model.Song
// @Failure		404	{object}	object
// @Failure		500	{object}	APIError
// @Router			/api/songs [get]
func (i *Implementation) GetSongs(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	songList, err := i.musicService.GetSongs(r.Context(), &model.Parameters{})
	if err != nil {
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	return i.JSON(w, http.StatusOK, songList)
}

// @Summary		Post Song
// @Description	Add song to Library
// @Tags			Songs
// @Accept			json
// @Produce		json
// @Param			song	body		request.AddSongRequest	true	"Add song"
// @Success		201		{object}	model.Song
// @Failure		404		{object}	object
// @Failure		422		{object}	APIError
// @Failure		500		{object}	APIError
// @Router			/api/songs [post]
func (i *Implementation) PostSong(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	var req request.AddSongRequest

	if err := decodeIntoStruct(r, &req); err != nil {
		if errors.Is(err, ErrNoBody) {
			return InvalidJSON()
		}
		return err
	}

	if errors := req.Validate(); errors != nil {
		return InvalidRequestData(errors)
	}

	song, err := i.musicService.PostSong(r.Context(), converter.FromAddSongRequestToModel(&req))
	if err != nil {
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	return i.JSON(w, http.StatusCreated, song)
}

// @Summary		Delete Song
// @Description	Deletes song from Library
// @Tags			Songs
// @Param			id	path	string	true	"Song ID"
// @Success		204
// @Failure		404	{object}	object
// @Failure		500	{object}	APIError
// @Router			/api/songs/{id} [delete]
func (i *Implementation) DeleteSong(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	var req request.ModifySongRequest

	req.ID = r.PathValue("id")

	if err := uuid.Validate(req.ID); err != nil {
		return NewAPIError(http.StatusBadRequest, ErrInvalidID)
	}

	if err := i.musicService.DeleteSong(r.Context(), converter.FromModifySongRequestToModel(&req)); err != nil {
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

// @Summary		Update Song Info
// @Description	Updates song information in Library
// @Tags			Songs
// @Accept			json
// @Produce		json
// @Param			id		path		string						true	"Song ID"
// @Param			song	body		request.ModifySongRequest	true	"Modify song info"
// @Success		200		{object}	model.Song
// @Failure		400		{object}	APIError
// @Failure		404		{object}	object
// @Failure		422		{object}	APIError
// @Failure		500		{object}	APIError
// @Router			/api/songs/{id} [put]
func (i *Implementation) UpdateSong(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	var req request.ModifySongRequest

	if err := decodeIntoStruct(r, &req); err != nil {
		if errors.Is(err, ErrNoBody) {
			return InvalidJSON()
		}
		return err
	}

	req.ID = r.PathValue("id")

	if err := uuid.Validate(req.ID); err != nil {
		return NewAPIError(http.StatusBadRequest, ErrInvalidID)
	}

	if errors := req.Validate(); errors != nil {
		return InvalidRequestData(errors)
	}

	song, err := i.musicService.UpdateSong(r.Context(), converter.FromModifySongRequestToModel(&req))
	if err != nil {
		if errors.Is(err, repo.ErrSongIsNotExist) {
			return NewAPIError(http.StatusBadRequest, err)
		}
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	return i.JSON(w, http.StatusOK, song)
}

func (i *Implementation) GetText(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	// TODO Implement

	i.log.DebugContext(r.Context(), "Handler done")

	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
