package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%+v\n", req)
	fmt.Fprintln(w, strings.Join(os.Environ(), "\n"))
	if req.URL.Path == "/crash" {
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	addr := ":" + port
	fmt.Printf("Listening on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
