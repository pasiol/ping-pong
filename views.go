package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	bytes, err := w.Write(response)
	if err != nil {
		log.Printf("writing response failed: %s", err)
	}
	log.Printf("response bytes %d", bytes)
}

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

func (a *App) getHealth(w http.ResponseWriter, _ *http.Request) {
	err := a.pool.Ping()
	if err == nil {
		respondWithJSON(w, http.StatusOK, "ok")
		return
	} else {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
