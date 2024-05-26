package main

import (
	"log"
	"net/http"
	"github.com/joshburnsxyz/go-wiki/pkg/api"
)

func main() {
	// Set up routes
	http.HandleFunc("/", api.FrontPageHandler)
	http.HandleFunc("/view/", api.MakeHandler(api.ViewHandler))
	http.HandleFunc("/edit/", api.MakeHandler(api.EditHandler))
	http.HandleFunc("/save/", api.MakeHandler(api.SaveHandler))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
