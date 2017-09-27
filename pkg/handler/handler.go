package handler

import (
	"net/http"

	"github.com/go-chi/chi"
)

type (
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
