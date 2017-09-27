package banana

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "github.com/makarski/mybanana/pkg/db/banana"
	"github.com/makarski/mybanana/pkg/handler"
	"github.com/makarski/mybanana/pkg/log"
)

type (
	// GetBananaHandler implements
	// http.Handler interface
	// and serves GET banana requests
	GetBananaHandler struct {
		finder      db.BananaFinder
		paramReader handler.URLParamReader
	}
)

// NewGetBananaHandler inits and returns an instance
// of GetBananaHandler
func NewGetBananaHandler(
	finder db.BananaFinder,
	paramReader handler.URLParamReader,
) http.Handler {
	return &GetBananaHandler{finder, paramReader}
}

// ServeHTTP implements http.Handler interface
func (h *GetBananaHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	errfmt := "GetBananaHandler: %s"

	bananaIDStr := h.paramReader.Read(req, "bananaID")
	bananaID, err := strconv.ParseUint(bananaIDStr, 10, 64)
	if err != nil {
		log.Error.Printf(errfmt, err)
		handler.Error(w, "invalid banana ID provided", http.StatusBadRequest)
		return
	}

	dbBanana, err := h.finder.Find(bananaID)
	if err != nil {
		log.Error.Printf(errfmt, err)
		handler.Error(w, "banana not found", http.StatusNotFound)
		return
	}

	banana := &Banana{}
	banana.fromDB(dbBanana)

	if err := json.NewEncoder(w).Encode([]*Banana{banana}); err != nil {
		log.Error.Printf(errfmt, err)
		handler.Error(w, "serialization error", http.StatusInternalServerError)
		return
	}
}
