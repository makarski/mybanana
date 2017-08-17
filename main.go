package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/makarski/mybanana/config"
	"github.com/makarski/mybanana/db"
	dbb "github.com/makarski/mybanana/db/banana"
	"github.com/makarski/mybanana/handler/banana"
	"github.com/makarski/mybanana/log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error.Fatal(err)
	}

	dbMap, err := db.NewDB(cfg.MySQLDSN, 5, 5, 0)
	if err != nil {
		log.Error.Fatal(err)
	}

	bananaFinder := dbb.NewBananaFinder(dbMap)

	router := chi.NewRouter()
	router.Route("/bananas", func(r chi.Router) {
		r.Get("/{bananaID:[0-9]+}", banana.NewGetBananaHandler(bananaFinder).ServeHTTP)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		log.Error.Fatal(err)
	}
}
