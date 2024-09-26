package rest

import (
	"errors"
	"net/http"

	"github.com/meraiku/music_lib/internal/api/rest/request"
	"github.com/meraiku/music_lib/internal/converter"
	"github.com/meraiku/music_lib/internal/model"
	"github.com/meraiku/music_lib/internal/repo"
	"go.uber.org/zap"
)

func statusCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("OK"))
}

func (i *Implementation) GetSongs(w http.ResponseWriter, r *http.Request) {

	i.log.DebugContext(r.Context(), "Handler started")

	songList, err := i.musicService.GetSongs(r.Context(), &model.Parameters{})
	if err != nil {
		i.log.Error("music service response",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	i.JSON(w, http.StatusOK, songList)
}

func (i *Implementation) PostSong(w http.ResponseWriter, r *http.Request) {

	var req request.AddSongRequest

	if err := decodeIntoStruct(r, &req); err != nil {
		if errors.Is(err, ErrNoBody) {
			w.WriteHeader(http.StatusBadRequest)
			i.ErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		i.log.Error("decoding struct",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := i.musicService.PostSong(r.Context(), converter.FromAddSongRequestToModel(&req)); err != nil {
		i.log.Error("posting song",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (i *Implementation) DeleteSong(w http.ResponseWriter, r *http.Request) {

	var req request.ModifySongRequest

	if err := decodeIntoStruct(r, &req); err != nil {
		i.log.Error("decoding struct",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.ID = r.PathValue("id")

	if req.ID == "" {
		i.ErrorJSON(w, http.StatusBadRequest, ErrNoBody.Error())
		return
	}

	if err := i.musicService.DeleteSong(r.Context(), converter.FromModifySongRequestToModel(&req)); err != nil {
		i.log.Error("deleting song",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (i *Implementation) UpdateSong(w http.ResponseWriter, r *http.Request) {

	var req request.ModifySongRequest

	if err := decodeIntoStruct(r, &req); err != nil {
		i.log.Error("decoding struct",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.ID == "" || req.Group == "" || req.Song == "" {
		i.ErrorJSON(w, http.StatusBadRequest, ErrNoBody.Error())
		return
	}

	if err := i.musicService.UpdateSong(r.Context(), converter.FromModifySongRequestToModel(&req)); err != nil {
		if errors.Is(err, repo.ErrSongIsNotExist) {
			i.ErrorJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		i.log.Error("updating song",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (i *Implementation) GetText(w http.ResponseWriter, r *http.Request) {

	// TODO Implement

	w.WriteHeader(http.StatusNotImplemented)
}
