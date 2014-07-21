package main

import (
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get("name")+"!\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
