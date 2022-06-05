package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	corsMiddleware "websitepreview.com/app/src/middleware/corsMiddleware"
	loggermiddleware "websitepreview.com/app/src/middleware/loggerMiddleware"
	webservice "websitepreview.com/app/src/services/webService"
)

func main() {

	port := "4000"
	router := mux.NewRouter()
	router.Use(corsMiddleware.CorsMiddleware)

	// https://github.com/anevsky/cachego

	// Controllers
	router.HandleFunc("/", loggermiddleware.LoggerMiddleware(webservice.HealthCheckHandler())).Methods("GET")
	router.HandleFunc("/scrape-link", webservice.ScrapeLink()).Methods("POST")
	router.HandleFunc("/scrape-link", webservice.ScrapeLink()).Methods("GET")
	router.HandleFunc("/session", webservice.GenerateSessionId()).Methods("GET")

	// Registering cron Jobs

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server is readt at", port)
	srv.ListenAndServe()
}
