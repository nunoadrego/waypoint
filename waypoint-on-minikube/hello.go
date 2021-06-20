package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting Hello Waypoint App.")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Waypoint!")
}
