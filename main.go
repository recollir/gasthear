package main

import (
	"io"
        "os"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
        name, _ := os.Hostname()
	io.WriteString(w, name+"\n")
}

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}	
