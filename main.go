package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO!"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show Snipper!"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Snippet!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", home)

	mux.HandleFunc("/snippet", showSnippet)

	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Server is running at port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
