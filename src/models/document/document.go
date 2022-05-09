package document

import (
	"net/http"
	"net/url"
)

type MetaTags struct {
	Description string `json:"description"`
}

type OgTags struct {
	OgType        string `json:"ogType"`
	OgTitle       string `json:"ogTitle"`
	OgDescription string `json:"ogDescription"`
	OgSiteName    string `json:"ogSiteName"`
	OgUrl         string `json:"ogUrl"`
	OgImage       string `json:"ogImage"`
}

type TwitterTags struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Card        string `json:"card"`
	Creator     string `json:"creator"`
	Domain      string `json:"domain"`
	Image       string `json:"image"`
}

type Document struct {
	Link            string         `json:"link"`
	TotalLinksCount int            `json:"totalLinksCount"`
	Heading         string         `json:"heading"`
	Language        string         `json:"language"`
	WordsCount      map[string]int `json:"wordsCount"`
	Title           string         `json:"title"`
	Links           []string       `json:"links"`
	Description     string         `json:"description"`
	Url             *url.URL       `json:"url"`
	Headers         *http.Header   `json:"headers"`
	Texts           []string       `json:"texts"`
	MetaTags        MetaTags       `json:"metaTags"`
	OgTags          OgTags         `json:"ogTags"`
	TwitterTags     TwitterTags    `json:"twitterTags"`
	LinkTexts       []string       `json:"linkTexts"`
	Canonical       string         `json:"canonical"`
	Keywords        []string       `json:"keywords"`
	Images          []string       `json:"images"`

	NumberOfInternalScripts int `json:"numberOfInternalScripts"`
	NumberOfExternalScripts int `json:"numberOfExternalScripts"`
	NumberOfOutboundLinks   int `json:"numberOfOutboundLinks"`
	NumberOfInboundLinks    int `json:"numberOfInboundLinks"`
	NumberOfExternalCss     int `json:"numberOfExternalCss"`
	NumberOfInternalCss     int `json:"numberOfInternalCss"`

	FavIcon  string `json:"favIcon"`
	Manifest string `json:"manifest"`

	HasViewportHtmlTag               bool `json:"hasViewportHtmlTag"`
	HasGoogleSiteVerificationHtmlTag bool `json:"hasGoogleSiteVerificationHtmlTag"`
}
