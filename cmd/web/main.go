package main

import (
	"github.com/orf1/basic-web-app-go/pkg/config"
	"github.com/orf1/basic-web-app-go/pkg/handlers"
	"github.com/orf1/basic-web-app-go/pkg/render"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template ceche")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Listening on port %s", port)
	_ = http.ListenAndServe(port, nil)
}
