package main

import (
	"net/http"
	"log"
)

func main() {
	// Serve static files from the "static" directory
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	fs := http.FileServer(http.Dir("."))
	// http.Handle("/", http.StripPrefix("/", fs))
	http.Handle("/", fs)



	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "index.html")
	// })

	port := ":3000"
	log.Println("Server running at http://localhost:3000/")
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
