package main

import (
	"net/http"

	redcap "github.com/tkruer/go-redcap/pkg"
)

func main() {
	client := redcap.RedCapClient{
		Token:          "1234567890",
		URL:            "https://redcap.example.edu/api/",
		ResponseFormat: "json",
	}

	http.HandleFunc("/get/arms", func(w http.ResponseWriter, r *http.Request) {
		client.ExportArms()
	})

	http.ListenAndServe("8080", nil)
}
