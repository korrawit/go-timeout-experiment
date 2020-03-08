package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Server received request")
		io.WriteString(w, "world")
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
