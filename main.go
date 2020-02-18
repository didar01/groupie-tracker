package main

import (
	"fmt"
	"log"
	"net/http"

	h "./static/go"
)

func main() {
	http.HandleFunc("/", h.Handler)
	http.HandleFunc("/Artist/", h.ArtistHandle)

	static := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", static))

	fmt.Println("listening on: http://localhost:2030")
	err := http.ListenAndServe(":2030", nil)

	if err != nil {
		log.Fatal(err)
	}

}
