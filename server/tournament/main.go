package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	tournamentAddr := os.Getenv("TOURNAMENTADDR")
	if len(tournamentAddr) == 0 {
		tournamentAddr = ":80"
	}
	mux := http.NewServeMux()
	//mux.HandleFunc("/v1/tournament", SummaryHandler)
	log.Printf("server is listening at %s...", tournamentAddr)
	log.Fatal(http.ListenAndServe(tournamentAddr, mux))
}
