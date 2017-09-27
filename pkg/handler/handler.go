package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type (
	apiError struct {
		Description string `json:"description"`
	}

	URLParamReader interface {
		Read(r *http.Request, name string) string
	}

	chiParamReader struct{}
)

func NewURLParamReader() URLParamReader {
	return &chiParamReader{}
}

func (*chiParamReader) Read(req *http.Request, name string) string {
	return chi.URLParam(req, name)
}

// Error writes an error to the provided writer
func Error(w http.ResponseWriter, err string, code int) {
	apiErr := apiError{err}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(apiErr)
}

// NotFound is a NotFound handler implementation
func NotFound(w http.ResponseWriter, req *http.Request) {
	Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
