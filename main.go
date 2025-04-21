package main

import (
	"fmt"
	"log"
	"net/http"
	"passive-reconnaissance/server"
)

func main() {
	app := server.NewApi()

	app.AddHandler("/hello", http.MethodGet, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from API!")
	})

	log.Println("Listening on :8080")

	if err := http.ListenAndServe(":8080", app); err != nil {
		return
	}
}
