package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Reading environment failed.")
	}
	a.DB, err = initializeDb()
	if err != nil {
		log.Fatalf("initializing database failed: %s", err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/pingpong", a.getPingPong).Methods("GET")
}

func (a *App) Run() {

	headers := handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{fmt.Sprintf("http://%s", os.Getenv("ALLOWED_ORIGINS"))})
	methods := handlers.AllowedMethods([]string{http.MethodGet, http.MethodOptions, http.MethodConnect, http.MethodPost})
	maxAge := handlers.MaxAge(60)

	address := fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT"))
	server := &http.Server{
		Addr:    address,
		Handler: handlers.CORS(headers, origins, methods, maxAge)(a.Router),
	}

	log.Printf("starting pingpong server in %s.", address)
	log.Printf("Version: %s , build: %s", Version, Build)
	log.Printf("Allowed origins: %s", os.Getenv("ALLOWED_ORIGINS"))
	log.Fatal(server.ListenAndServe())
}
