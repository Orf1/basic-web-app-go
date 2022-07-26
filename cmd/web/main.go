package main

import (
	"fmt"
	"github.com/orf1/basic-web-app-go/pkg/handlers"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Println(fmt.Sprintf("Listening on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
