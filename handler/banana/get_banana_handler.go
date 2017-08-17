package banana

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	db "github.com/makarski/mybanana/db/banana"
	"github.com/makarski/mybanana/log"
)

type (
	// GetBananaHandler implements
	// http.Handler interface
	// and serves GET banana requests
	GetBananaHandler struct {
		finder db.BananaFinder
	}
)

// NewGetBananaHandler inits and returns an instance
// of GetBananaHandler
func NewGetBananaHandler(finder db.BananaFinder) http.Handler {
	return &GetBananaHandler{finder}
}

// ServeHTTP implements http.Handler interface
func (h *GetBananaHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	errfmt := "GetBananaHandler: %s"

	bananaIDStr := chi.URLParam(req, "bananaID")
	bananaID, err := strconv.ParseUint(bananaIDStr, 10, 64)
	if err != nil {
		log.Error.Printf(errfmt, err)
		http.Error(w, "invalid banana ID provided", http.StatusBadRequest)
		return
	}

	dbBanana, err := h.finder.Find(bananaID)
	if err != nil {
		log.Error.Printf(errfmt, err)
		http.Error(w, "banana not found", http.StatusNotFound)
		return
	}

	banana := &Banana{}
	banana.fromDB(dbBanana)

	b, err := json.Marshal([]*Banana{banana})
	if err != nil {
		log.Error.Printf(errfmt, err)
		http.Error(w, "serialization error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
