package main

import (
	"fmt"
	"net/http"
	"os"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is the server responding!")
}

func main() {
	fmt.Printf("kcd-server running with pid: %d", os.Getpid())
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}
