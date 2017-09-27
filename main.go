package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/makarski/mybanana/pkg/config"
	"github.com/makarski/mybanana/pkg/db"
	dbb "github.com/makarski/mybanana/pkg/db/banana"
	"github.com/makarski/mybanana/pkg/handler"
	"github.com/makarski/mybanana/pkg/handler/banana"
	"github.com/makarski/mybanana/pkg/log"
	"github.com/makarski/mybanana/pkg/middleware"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error.Panic(err)
	}

	dbMap, err := db.NewDB(cfg.MySQLDSN, 5, 5, 0)
	if err != nil {
		log.Error.Fatal(err)
	}

	bananaFinder := dbb.NewBananaFinder(dbMap)
	urlReader := handler.NewURLParamReader()

	router := chi.NewRouter()
	// We want to have a unified JSON response
	// for Not Found responses as well
	router.NotFound(handler.NotFound)

	// We set middlewares we want to use
	router.Use(
		middleware.ContentTypeJSON,
	)

	// We define our routes
	router.Route("/bananas", func(r chi.Router) {
		r.Get("/{bananaID:[0-9]+}", banana.NewGetBananaHandler(bananaFinder, urlReader).ServeHTTP)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		log.Error.Fatal(err)
	}
}
