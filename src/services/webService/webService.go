package webservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	scrapelinkpayload "websitepreview.com/app/src/payloads/scrapeLinkPayload"
	scrapeservice "websitepreview.com/app/src/services/scrapingService"
)

func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"status": 200})
	}
}

func GenerateSessionId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		uuidStr := uuid.String()
		w.WriteHeader(200)
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"sessionId": uuidStr})
	}
}

func ScrapeLink() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")

		sessionId := r.Header.Get("Authorization")
		fmt.Println(sessionId)

		if r.Method == "POST" {
			var p scrapelinkpayload.ScrapeLinkPayload
			err := json.NewDecoder(r.Body).Decode(&p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			document := scrapeservice.ScrapeLink(p.Link)
			json.NewEncoder(w).Encode(document)

		} else if r.Method == "GET" {
			link := r.URL.Query().Get("link")
			document := scrapeservice.ScrapeLink(link)
			json.NewEncoder(w).Encode(document)
		}
	}
}

func ScrapeDomain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		var p scrapelinkpayload.ScrapeLinkPayload

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		document := scrapeservice.ScrapeLink(p.Link)
		// links := scrapeservice.ScrapeFirstPageLinks(p.Link)

		// TODO. Save Into DB
		json.NewEncoder(w).Encode(document)
	}
}
