package main

import (
	"net/http"
)

func main() {
	servemux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: servemux,
	}
	servemux.Handle("/", http.FileServer(http.Dir(".")))
	server.ListenAndServe()
}
