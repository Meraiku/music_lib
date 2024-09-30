package rest

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
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

func decodeIntoStruct(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&v)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return InvalidJSON()
		}
		return fmt.Errorf("error decoding parameters: %s", err)
	}

	return nil
}
