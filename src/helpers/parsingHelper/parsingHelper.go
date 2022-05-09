package parsinghelper

import (
	"strings"

	"github.com/gocolly/colly"
)

func HasKnownBrandDomain(link string) bool {
	if strings.Contains(link, "google.com") ||
		strings.Contains(link, "twitter.com") ||
		strings.Contains(link, "facebook.com") ||
		strings.Contains(link, "apple.com") ||
		strings.Contains(link, "instagram.com") {
		return true
	} else {
		return false
	}
}

func RemoveLinksWithKnownBrandDomain(links []string) []string {
	lnks := make([]string, 0)
	for _, link := range links {
		includesTargetBrand := HasKnownBrandDomain(link)
		if !includesTargetBrand {
			lnks = append(lnks, link)
		}
	}

	return lnks
}

func IsFullLink(link string) bool {
	if strings.HasPrefix(link, "https") || strings.HasPrefix(link, "http") {
		return true
	} else {
		return false
	}
}

func TransformLinks(links []string, domain string) []string {
	cleanLinks := RemoveLinksWithKnownBrandDomain(links)
	lnks := make([]string, 0)

	for _, link := range cleanLinks {
		isFull := IsFullLink(link)
		if isFull {
			lnks = append(lnks, link)
		} else {
			lnks = append(lnks, domain+link)
		}
	}

	return lnks
}

func ConstructLink(rawLink string, domainWithProtocol string) (link string) {
	if IsFullLink(rawLink) {
		return rawLink
	}

	// Construct array with single link and remove first / since domain.url includes last slash
	linksArr := make([]string, 0)
	linksArr = append(linksArr, rawLink)

	constructedLinks := TransformLinks(linksArr, domainWithProtocol)
	if len(constructedLinks) == 0 {
		return ""
	}
	return constructedLinks[0]
}

func GetWordsCount(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, w := range words {
		value, ok := wordCounts[w]
		if ok {
			wordCounts[w] = value + 1
		} else {
			wordCounts[w] = 1
		}
	}

	return wordCounts
}

func CountNumberOfExternalScripts(numberOfExternalScripts *int, elem *colly.HTMLElement) {
	if strings.HasPrefix(elem.Attr("src"), "https") || strings.HasPrefix(elem.Attr("src"), "http") {
		*numberOfExternalScripts = *numberOfExternalScripts + 1
	}
}

func CountNumberOfInternalScripts(numberOfInternalScripts *int, elem *colly.HTMLElement) {
	// Basic check
	// External: https://cdn.com/app.js
	// Internal: /app.js
	if strings.HasPrefix(elem.Attr("src"), "/") || len(elem.Attr("src")) == 0 {
		*numberOfInternalScripts = *numberOfInternalScripts + 1
	}
}

func CountNumberOfInternalCss(numberOfInternalCss *int, elem *colly.HTMLElement) {
	if strings.HasPrefix(elem.Attr("href"), "/") || len(elem.Attr("href")) == 0 {
		*numberOfInternalCss = *numberOfInternalCss + 1
	}
}

func CountNumberOfExternalCss(numberOfExternalCss *int, elem *colly.HTMLElement) {
	if strings.HasPrefix(elem.Attr("href"), "https") || strings.HasPrefix(elem.Attr("href"), "http") {
		*numberOfExternalCss = *numberOfExternalCss + 1
	}
}

func CountNumberOfInboundLinks(numberOfInboundLinks *int, elem *colly.HTMLElement) {
	if strings.HasPrefix(elem.Attr("href"), "/") || len(elem.Attr("href")) == 0 {
		*numberOfInboundLinks = *numberOfInboundLinks + 1
	}
}

func CountNumberOfOutboundLinks(numberOfOutboundLinks *int, elem *colly.HTMLElement) {
	if strings.HasPrefix(elem.Attr("href"), "https") || strings.HasPrefix(elem.Attr("href"), "http") {
		*numberOfOutboundLinks = *numberOfOutboundLinks + 1
	}
}
