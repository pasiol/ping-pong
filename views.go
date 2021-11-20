package main

import (
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
)

func (a *App) getPingPong(w http.ResponseWriter, r *http.Request) {
	var last PingPong
	result := a.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Last(&last)
	if result.Error != nil {
		log.Printf("getting counter failed: %s", result.Error)
	}
	last.Counter++
	counter := last.Counter
	result = a.DB.Save(&last)
	if result.Error != nil {
		log.Printf("updating counter failed: %s", result.Error)
	}
	output := fmt.Sprintf("Ping / Pongs: %d\n", counter)
	b, err := fmt.Fprintf(w, output)
	if err != nil {
		log.Fatalf("writing response failed %s", r.RemoteAddr)
	}
	if len(output) > 0 {
		log.Printf("written %d bytes address %s: %s", b, r.RemoteAddr, output[0:len(output)-1])
	}
}
