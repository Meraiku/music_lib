package rest

import (
	"errors"
	"net/http"

	"github.com/meraiku/music_lib/internal/api/rest/request"
	"github.com/meraiku/music_lib/internal/converter"
	"github.com/meraiku/music_lib/internal/model"
	"github.com/meraiku/music_lib/internal/repo"
)

func statusCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("OK"))
}

func (i *Implementation) GetSongs(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	songList, err := i.musicService.GetSongs(r.Context(), &model.Parameters{})
	if err != nil {
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	return i.JSON(w, http.StatusOK, songList)
}

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

	if err := i.musicService.PostSong(r.Context(), converter.FromAddSongRequestToModel(&req)); err != nil {
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	w.WriteHeader(http.StatusCreated)
	return nil
}

func (i *Implementation) DeleteSong(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	var req request.ModifySongRequest

	req.ID = r.PathValue("id")

	if req.ID == "" {
		i.ErrorJSON(w, http.StatusBadRequest, ErrNoBody.Error())
		return nil
	}

	if err := i.musicService.DeleteSong(r.Context(), converter.FromModifySongRequestToModel(&req)); err != nil {
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	w.WriteHeader(http.StatusNoContent)
	return nil
}

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

	if errors := req.Validate(); errors != nil {
		return InvalidRequestData(errors)
	}

	if err := i.musicService.UpdateSong(r.Context(), converter.FromModifySongRequestToModel(&req)); err != nil {
		if errors.Is(err, repo.ErrSongIsNotExist) {
			return NewAPIError(http.StatusBadRequest, err)
		}
		return err
	}

	i.log.DebugContext(r.Context(), "Handler done")

	w.WriteHeader(http.StatusOK)
	return nil
}

func (i *Implementation) GetText(w http.ResponseWriter, r *http.Request) error {

	i.log.DebugContext(r.Context(), "Handler started")

	// TODO Implement

	i.log.DebugContext(r.Context(), "Handler done")

	w.WriteHeader(http.StatusNotImplemented)
	return nil
}
