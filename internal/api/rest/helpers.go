package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

func (i *Implementation) JSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(&payload)
	if err != nil {
		i.log.Error("marshaling json",
			zap.Error(err),
		)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(code)
	w.Write(data)
}

func (i *Implementation) ErrorJSON(w http.ResponseWriter, code int, msg string) {
	type ErrMsg struct {
		Err string `json:"error"`
	}

	resp := &ErrMsg{
		Err: msg,
	}

	i.JSON(w, code, resp)
}

func decodeIntoStruct(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(v)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return ErrNoBody
		}
		return fmt.Errorf("error decoding parameters: %s", err)
	}

	return nil
}
