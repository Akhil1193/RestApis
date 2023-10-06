package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &home{})

	http.ListenAndServe(":8080", mux)
}

// type RecipesHandler struct{}

// func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Dont know how to cook"))
// }

type home struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is my home page"))
}
