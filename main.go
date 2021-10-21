package main

import (
	"fmt"
	"net/http"
)

func print(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This App is Running")
	fmt.Fprint(w,"Hook is triggered")  // Adding comment to check git hooks
}

func main() {
	http.HandleFunc("/", print)
	http.ListenAndServe(":8085", nil)
}
