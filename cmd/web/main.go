package main

import (
	"fmt"
	"github.com/GoloisaNinja/go-hello_world/pkg/config"
	"github.com/GoloisaNinja/go-hello_world/pkg/handlers"
	"github.com/GoloisaNinja/go-hello_world/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const PORT = ":8000"

// boolean switch that impacts session cookie secure in both scs and middleware
var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction
	
	app.Session = session

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
