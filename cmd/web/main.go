package main

import (
	"fmt"
	"github.com/GoloisaNinja/go-hello_world/pkg/handlers"
	"log"
	"net/http"
)

const PORT = ":8000"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// serve up the css

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Printf("Server is up on port %s", PORT)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
