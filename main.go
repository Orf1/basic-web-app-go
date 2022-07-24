package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")

		log.Println(fmt.Sprintf("Number of bytes written: %d", n))

		if err != nil {
			log.Println(fmt.Sprintf("An error occurred: %d", err))
		}
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Println(fmt.Sprintf("An error occurred: %d", err))
	}
}
