package main

import (
	"net/http"
	"os"

	"github.com/kenjitheman/seadclub_bot/tg"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	go tg.Start()

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
