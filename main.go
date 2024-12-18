package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve the HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// HTMX endpoint
	http.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<p class="text-green-600">This content was loaded dynamically with HTMX!</p>`)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
