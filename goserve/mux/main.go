package main

import (
	"net/http"
)

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", someText)
	http.ListenAndServe(":8080", myMux)
}
func someText(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("This is a mux"))
}
