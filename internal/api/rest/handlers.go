package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/meraiku/music_lib/internal/api/rest/request"
	"github.com/meraiku/music_lib/internal/converter"
	"github.com/meraiku/music_lib/internal/lib/fetcher"
	"github.com/meraiku/music_lib/internal/model"
	"github.com/meraiku/music_lib/internal/repo"
)

// @Summary	Check Server Availability
// @Produce	json
// @Success	200	{object}	object
// @Failure	404	{object}	object
// @Router		/healthz [get]
func (i *Implementation) ServerStatus(w http.ResponseWriter, r *http.Request) {
	i.JSON(w, http.StatusOK, struct{}{})
}

// @Summary		Get Songs
// @Description	Prints List of songs from library.
// @Tags			Songs
// @Produce		json
// @Param			page	query		int		false	"Page number. Default 1"
// @Param			filter	query		string	false	"Filter By {}. Default 'song' name"
// @Success		200		{array}		model.Song
// @Failure		404		{object}	object
// @Failure		500		{object}	APIError
// @Router			/api/songs [get]
func (i *Implementation) GetSongs(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	var p model.Parameters
	var err error

	query := r.URL.Query()

	p.Page, _ = strconv.Atoi(query.Get("page"))
	p.Filter = query.Get("filter")

	if p.Filter == "" {
		p.Filter = "song"
	}
	if p.Page <= 0 {
		p.Page = 1
	}

	if err := p.Validate(); err != nil {
		return NewAPIError(http.StatusBadRequest, err)
	}

	fmt.Println(p.Page)

	songList, err := i.musicService.GetSongs(r.Context(), &p)
	if err != nil {
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	return i.JSON(w, http.StatusOK, songList)
}

// @Summary		Get Song Text
// @Description	Prints text with verse number
// @Tags			Songs
// @Produce		json
// @Param			verse	query		string	false "Verse number"
// @Param			id		path		string	true	"Song ID"
// @Success		200		{array}		model.Text
// @Failure		400		{object}	APIError
// @Failure		404		{object}	object
// @Failure		422		{object}	APIError
// @Failure		500		{object}	APIError
// @Router			/api/songs/{id}/text [get]
func (i *Implementation) GetText(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	id := r.PathValue("id")

	if err := uuid.Validate(id); err != nil {
		return NewAPIError(http.StatusBadRequest, ErrInvalidID)
	}

	query := r.URL.Query()

	verseNumber, _ := strconv.Atoi(query.Get("verse"))

	text, err := i.musicService.GetText(r.Context(), id, verseNumber)
	if err != nil {
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	return i.JSON(w, http.StatusOK, text)
}

// @Summary		Post Song
// @Description	Enriches song with additional inforamtion, then adds song to Library. If song inforamtion can't be enriched, error is shown
// @Tags			Songs
// @Accept			json
// @Produce		json
// @Param			song	body		request.AddSongRequest	true	"Band and Song names"
// @Success		201		{object}	model.Song
// @Failure		400		{object}	APIError
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
		if errors.Is(err, fetcher.ErrNoData) {
			return NewAPIError(http.StatusBadRequest, err)
		}
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

	id := r.PathValue("id")

	if err := uuid.Validate(id); err != nil {
		return NewAPIError(http.StatusBadRequest, ErrInvalidID)
	}

	if err := i.musicService.DeleteSong(r.Context(), id); err != nil {
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
// @Router			/api/songs/{id} [patch]
func (i *Implementation) UpdateSong(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	var req request.PatchRequest

	if err := decodeIntoStruct(r, &req); err != nil {
		if errors.Is(err, ErrNoBody) {
			return InvalidJSON()
		}
		return err
	}

	id := r.PathValue("id")

	if err := uuid.Validate(id); err != nil {
		return NewAPIError(http.StatusBadRequest, ErrInvalidID)
	}

	song, err := i.musicService.UpdateSong(r.Context(), converter.FromPatchRequestToModel(&req, id))
	if err != nil {
		if errors.Is(err, repo.ErrSongIsNotExist) {
			return NewAPIError(http.StatusBadRequest, err)
		}
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	return i.JSON(w, http.StatusOK, song)
}
