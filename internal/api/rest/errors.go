package rest

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

var (
	ErrNoBody = errors.New("request body is empty")
)

type APIError struct {
	StatusCode int `json:"status_code"`
	Msg        any `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("api error: %s", e.Msg)
}

func NewAPIError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Msg:        err.Error(),
	}
}

func InvalidRequestData(errors map[string]string) APIError {
	return APIError{
		StatusCode: http.StatusUnprocessableEntity,
		Msg:        errors,
	}
}

func InvalidJSON() APIError {
	return NewAPIError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"))
}

type APIFunc func(w http.ResponseWriter, r *http.Request) error

func (i *Implementation) Make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {

			if apiErr, ok := err.(APIError); ok {
				i.JSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"status_code": http.StatusInternalServerError,
					"msg":         "internal server error",
				}
				i.JSON(w, http.StatusInternalServerError, errResp)
			}
			i.log.ErrorContext(r.Context(), "HTTP API error", slog.Any("error", err))
		}
	}
}
