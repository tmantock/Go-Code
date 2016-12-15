package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hellomars)
	http.ListenAndServe(":8080", nil)
}

func hellomars(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello Mars!"))
}
