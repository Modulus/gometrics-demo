package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	mainRouter := mux.NewRouter()
	mainRouter.HandleFunc("/", RootHandler).Methods("GET")
	mainRouter.HandleFunc("/vote/{category}", VoteHandler).Methods("GET")
	log.Println("Server running on 0.0.0.0:8000")

	mainHttp := &http.Server{Addr: ":8000", Handler: mainRouter}

	metricsRouter := mux.NewRouter()
	metricsRouter.Path("/metrics").Handler(promhttp.Handler())
	metricsHttp := &http.Server{Addr: ":2112", Handler: metricsRouter}

	go func() {
		log.Println("Server running on 0.0.0.0:8000")
		if err := mainHttp.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Main HTTP server closed")
			} else {
				log.Panic("Could not start ai-api-gateway http server", err)
			}
		}
	}()

	go func() {
		log.Println("Metrics Server running on 0.0.0.0:2112")
		if err := metricsHttp.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Metrics HTTP server closed")
			} else {
				log.Panic("Could not start ai-api-gateway http server", err)
			}
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down the Gateway server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	mainHttp.Shutdown(ctx)
	metricsHttp.Shutdown(ctx)

	log.Println("App exiting")

}
