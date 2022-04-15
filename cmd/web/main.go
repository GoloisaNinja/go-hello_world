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
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// serve up the css

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Server is up on port %s", PORT)

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
