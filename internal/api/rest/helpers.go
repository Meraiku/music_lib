package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (i *Implementation) JSON(w http.ResponseWriter, code int, payload any) error {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(&payload)
	if err != nil {
		return err
	}

	w.WriteHeader(code)
	_, err = w.Write(data)
	return err
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
