package webservice

import (
	"encoding/json"
	"net/http"

	scrapelinkpayload "websitepreview.com/app/src/payloads/scrapeLinkPayload"
	scrapeservice "websitepreview.com/app/src/services/scrapingService"
)

func HealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"status": 200})
	}
}

func ScrapeLink() http.HandlerFunc {
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
		json.NewEncoder(w).Encode(document)
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
