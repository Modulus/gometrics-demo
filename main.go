package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Data struct {
	TimeStamp string
	Name      string
	Message   string
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	log.Println("You voted for " + category)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	data := Data{currentTime, "You voted for", category}
	js, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Server", "Ein slek der Mix Master")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	data := Data{currentTime, "First post", "This is the first post"}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Server", "Ein slek der Mix Master")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/vote/{category}", VoteHandler)
	log.Println("Server running on 0.0.0.0:8000")

	r.Path("/metrics").Handler(promhttp.Handler())

	http.ListenAndServe(":8000", r)
}
