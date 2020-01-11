package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default to port %s", port)
	}

	log.Printf("Listening on port: %s", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
