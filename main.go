package main

import (
	"log"
	"net/http"
	"os"

	"github.com/FreshworksStudio/bs-go-utils/lib"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/start", start)
	http.HandleFunc("/move", move)
	http.HandleFunc("/end", end)
	http.HandleFunc("/ping", ping)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Running server on port %s...\n", port)
	http.ListenAndServe(":"+port, lib.LoggingHandler(http.DefaultServeMux))
}
