package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	loggermiddleware "websitepreview.com/app/src/middleware/loggerMiddleware"
	webservice "websitepreview.com/app/src/services/webService"
)

func main() {

	port := "5300"

	fmt.Println("Registering routes .....")
	router := mux.NewRouter()

	// stopper := scraperjob.ScrapeDomainsJob()
	// defer close(stopper)

	// Controllers
	router.HandleFunc("/", loggermiddleware.LoggerMiddleware(webservice.HealthCheckHandler())).Methods("GET")
	router.HandleFunc("/scrape-link", webservice.ScrapeLink()).Methods("POST")

	// Registering cron Jobs

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	srv := &http.Server{
		Handler: handler,
		Addr:    "127.0.0.1:" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server is readt at", port)
	srv.ListenAndServe()
}
