package scraperjob

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	scrapeservice "websitepreview.com/app/src/services/scrapingService"
	"websitepreview.com/app/src/utils/scrapeutils"
)

func ScrapeDomainsJob() chan bool {
	domains := scrapeutils.GetDomains()
	s := gocron.NewScheduler()
	s.Every(15).Seconds().Do(func() {
		for _, domain := range domains {
			document := scrapeservice.ScrapeLink(domain.Link)
			fmt.Println(document.Title)
		}
	})
	stopper := s.Start()
	return stopper
}
