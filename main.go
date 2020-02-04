package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", messageHandler) // Path to Route
	http.ListenAndServe(":8080", nil)    // Port

}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Let's get to building!")
}
