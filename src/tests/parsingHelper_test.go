package tests

import (
	"testing"

	parsinghelper "websitepreview.com/app/src/helpers/parsingHelper"
)

func TestKnownDomainDetectorHelper(t *testing.T) {

	domain := "google.com"
	hasKnownBrandName := parsinghelper.HasKnownBrandDomain(domain)
	if !hasKnownBrandName {
		t.Errorf("Failed to detect known domain!")
	}
}

func TestRemoveLinksWithKnownBrandDomainHelper(t *testing.T) {
	links := make([]string, 0)
	links = append(links, "https://google.com")
	links = append(links, "https://facebook.com")
	links = append(links, "https://mysite.com")

	links = parsinghelper.RemoveLinksWithKnownBrandDomain(links)
	if links[0] != "https://mysite.com" {
		t.Errorf("Failed to remove known brand domain")
	}
}

func TestTransformLinksHelper(t *testing.T) {
	domain := "https://rustavi2.ge"
	links := make([]string, 0)
	links = append(links, "/ka/test-page")
	links = append(links, "https://rustavi2.ge/test-page2")

	expectedFirstLink := "https://rustavi2.ge/ka/test-page"
	links = parsinghelper.TransformLinks(links, domain)
	if links[0] != expectedFirstLink {
		t.Errorf("Invalid Link")
	}

}
