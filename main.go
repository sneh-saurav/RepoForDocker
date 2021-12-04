package main

import (
	"fmt"
	"net/http"
)

func print(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This App is Running")
	fmt.Fprintln(w," I hope Hook is been triggered")   // New comment again
}

func main() {
	http.HandleFunc("/", print)
	http.ListenAndServe(":8085", nil)
}
