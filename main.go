package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var counter = 0

func createDir(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("creating directory failed: %s", err)
	}
}

func createFile(filename string, content string) {
	if !fileExists(filepath.Dir(filename)) {
		createDir(filepath.Dir(filename))
	}
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatalf("creating file %s failed: %s", filename, err)
	}
}

func write2File(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("opening file %s failed: %s", filename, err)
	}
	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatalf("writing to file %s failed: %s", filename, err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("closing file %s failed: %s", filename, err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	output := fmt.Sprintf("Ping / Pongs: %d\n", counter)
	_, err := fmt.Fprintf(w, output)
	if err != nil {
		log.Fatalf("writing response failed %s", r.RemoteAddr)
	}
	log.Printf(output)
	if !fileExists(os.Getenv("APP_OUTPUT_FILE")) {
		createFile(os.Getenv("APP_OUTPUT_FILE"), output)
	}
	write2File(os.Getenv("APP_OUTPUT_FILE"), output)
}

func main() {
	if !fileExists(os.Getenv("APP_LOG_FILE")) {
		createFile(os.Getenv("APP_LOG_FILE"), "")
	}
	f, err := os.OpenFile(os.Getenv("APP_LOG_FILE"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("opening log file failed: %s", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("closing log file failed: %s", err)
		}
	}(f)

	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	port := os.Getenv("APP_PORT")

	address := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("pingopong starting in port %s.", address)
	http.HandleFunc("/pingpong", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
