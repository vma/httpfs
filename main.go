package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/vma/httplogger"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	logger := httplogger.CombinedLogger(os.Stdout)
	http.Handle("/", http.FileServer(http.Dir(*directory)))
	log.Printf("Serving %s over http port: %s", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, logger(http.DefaultServeMux)))
}
