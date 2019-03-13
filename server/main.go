package main

import (
	"net/http"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")

	mux := http.NewServeMux()
	http.ListenAndServe(addr, mux)
}
