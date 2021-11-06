package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var counter = 0

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	output := fmt.Sprintf("Ping / Pongs: %d", counter)
	_, err := fmt.Fprintf(w, output)
	if err != nil {
		log.Fatalf("writing response failed %s", r.RemoteAddr)
	}
	log.Printf(output)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	address := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("go-pingopong starting in port %s.", address)
	http.HandleFunc("/pingpong", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
