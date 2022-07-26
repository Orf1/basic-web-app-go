package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Println(fmt.Sprintf("Listening on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
