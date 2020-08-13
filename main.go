package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/a5a5a9/api-calculator/handlers"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var serverPort string = "8080"
var started = time.Now()

func main() {
	var log = logrus.New()
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.WelcomeHandler)
	r.HandleFunc("/add/{numbers}", handlers.AddHandler)
	r.HandleFunc("/substract/{numbers}", handlers.SubstractHandler)
	r.HandleFunc("/division/{numbers}", handlers.DivisionHandler)
	r.HandleFunc("/random/{count}", handlers.RandomHandler)
	r.HandleFunc("/random", handlers.RandomHandler)
	r.HandleFunc("/liveness", handlers.LivenessCheckHandler)
	r.HandleFunc("/readiness", handlers.ReadinessCheckHandler)
	// exposes /metrics endpoint with standard golang metrics used by prometheus
	r.Handle("/metrics", promhttp.Handler())

	fmt.Printf("Webservice is up on localhost:" + serverPort)
	http.ListenAndServe(":"+serverPort, r)
	// / Bind to a port and pass our loggedRouter in
	// log.Fatal(http.ListenAndServe(":8080", metricsWrapper))
}
