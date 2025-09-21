package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
// 	maxBytes := 1048576 //one megabayt

// 	fmt.Println("r.body", r.Body)
// 	fmt.Println("", data)
// 	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

// 	dec := json.NewDecoder(r.Body)
// 	fmt.Println("dec", dec)
// 	err := dec.Decode(data)
// 	fmt.Println("data decode sonrası", data)
// 	if err != nil {
// 		fmt.Println("decode hatası")
// 		return err
// 	}

// 	err = dec.Decode(&struct{}{})
// 	if err != io.EOF {
// 		fmt.Println("okuma hatası")

// 		return errors.New("body must have only a single value")
// 	}

// 	return nil
// }

func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, v any) error {
	const maxBytes = 1 << 20 // 1 MB

	// Body boyutunu sınırla
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields() // (opsiyonel) JSON’daki tanımsız alanları yakalar

	// Gelen JSON'u v parametresine decode et
	if err := dec.Decode(v); err != nil {
		fmt.Println("decode hatası:", err)
		return err
	}

	// Fazladan veri olup olmadığını kontrol et
	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must contain only a single JSON object")
	}

	return nil
}

func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {

		return err
	}

	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)
}
