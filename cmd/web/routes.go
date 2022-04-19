package main

import (
	"github.com/GoloisaNinja/go-hello_world/pkg/config"
	"github.com/GoloisaNinja/go-hello_world/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	// example...
	//mux.Use(WriteToConsole)

	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	// load the css
	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	return mux
}
