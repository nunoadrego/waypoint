package hello

import (
	"fmt"
	"net/http"
)

func Hello() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Waypoint!")
}
