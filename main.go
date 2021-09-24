package main

import (
	"fmt"
	"net/http"
)

func print(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This App is Running")
}

func main() {
	http.HandleFunc("/", print)
	http.ListenAndServe(":8085", nil)
}