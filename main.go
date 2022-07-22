package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Println("Listening on Port 3000...")
	http.ListenAndServe(":3000", nil)
}
