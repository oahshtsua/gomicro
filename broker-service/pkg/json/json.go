package json

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Read(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1048576 // 1mb

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dst)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single json")
	}
	return nil
}

func Write(w http.ResponseWriter, status int, data any, headers http.Header) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.WriteHeader(status)
	w.Write(js)
	return nil
}

func Error(w http.ResponseWriter, err error, status int) error {
	resp := Response{
		Error:   true,
		Message: err.Error(),
	}
	return Write(w, status, resp, nil)
}
