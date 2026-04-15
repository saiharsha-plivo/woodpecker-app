package main 

import (
	"net/http"
	"log"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	log.Println("Starting server on port 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error starting server", err)
		log.Fatal(err)
	}
}