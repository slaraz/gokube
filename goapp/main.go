package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	hn, _ := os.Hostname()
	fmt.Fprintf(w, "%q: Cześć World! v4 %s", hn, time.Now())
}

func main() {
	log.Println("Start.")
	defer log.Println("Koniec.")

	http.HandleFunc("/", greet)
	http.ListenAndServe(":http", loguj(http.DefaultServeMux))
}

func loguj(wewn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wewn.ServeHTTP(w, r)
		log.Print(r.RequestURI, time.Since(start))
	})
}
