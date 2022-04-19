package main

import (
	"fmt"
	"github.com/GoloisaNinja/go-hello_world/pkg/config"
	"github.com/GoloisaNinja/go-hello_world/pkg/handlers"
	"github.com/GoloisaNinja/go-hello_world/pkg/render"
	"log"
	"net/http"
)

const PORT = ":8000"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache...")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Server is up on port %s", PORT)

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
