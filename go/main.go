package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func checkCOR(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originHeader := r.Header.Get("Origin")
		originUrl, _ := url.Parse(originHeader)
		if strings.HasSuffix(originUrl.Hostname(), ".legalsink.com") {
			w.Header().Add("Access-Control-Allow-Origin", originHeader)
		}
		h.ServeHTTP(w, r)
	})
}

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

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/resource", resourceHandler)
	mux.HandleFunc("/docs", docsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default to port %s", port)
	}

	log.Printf("Listening on port: %s", port)

	log.Fatal(http.ListenAndServe(":"+port, checkCOR(mux)))
}
