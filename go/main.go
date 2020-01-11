package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello\n")
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Getting Resource\n")
}

func docsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Text for Docs\n")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/resource", resourceHandler)
	http.HandleFunc("/docs", docsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
