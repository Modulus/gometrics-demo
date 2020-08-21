package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/log"
)

type Data struct {
	Id      int
	Name    string
	Message string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	data := Data{1024, "First post", "This is the first post"}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Server", "Ein slek der Mix Master")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)

	log.Info(http.ListenAndServe(":8000", r))
}
