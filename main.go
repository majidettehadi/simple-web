package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var startTime time.Time

func main () {
	startTime = time.Now()
	router := mux.NewRouter()
	router.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handler (w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	response := "I'm Up (hostname: "+ hostname +" - uptime: " + shortDur(time.Since(startTime)) + " )"
	w.WriteHeader(200)
	_, _ = w.Write([]byte(response))

}

func shortDur(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
