package scrapeservice

import (
	"strings"
	"time"

	"github.com/gocolly/colly"
	"websitepreview.com/app/src/helpers"
	parsinghelper "websitepreview.com/app/src/helpers/parsingHelper"
	"websitepreview.com/app/src/models/document"
)

func ScrapeFirstPageLinks(domain string) []string {
	c := colly.NewCollector(colly.MaxDepth(1))

	document := new(document.Document)
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		// r.Ctx.Put("url", r.URL.String())
		r.Headers.Set("User-Agent", helpers.RandomUser())
		// document.Url = r.URL
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		document.Links = append(document.Links, e.Attr("href"))
	})

	c.Visit(domain)

	transformedLinks := parsinghelper.TransformLinks(document.Links, domain)
	return transformedLinks
}

func ScrapeLink(link string) document.Document {

	c := colly.NewCollector(colly.MaxDepth(1))

	document := new(document.Document)
	document.Link = link

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay:       5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		r.Ctx.Put("url", r.URL.String())
		r.Headers.Set("User-Agent", helpers.RandomUser())
		document.Url = r.URL
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		document.TotalLinksCount = document.TotalLinksCount + 1
		link := parsinghelper.ConstructLink(e.Attr("href"), document.Url.String())
		if link == "" {
			return
		}
		document.Links = append(document.Links, link)
	})

	c.OnResponse(func(r *colly.Response) {
		document.Headers = r.Headers
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		document.Title = e.Text
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		document.Heading = e.Text
	})

	c.OnHTML("description", func(e *colly.HTMLElement) {
		document.Description = e.Text
	})

	c.OnHTML("html", func(e *colly.HTMLElement) {

		texts := make([]string, 0)
		words := make([]string, 0)
		images := make([]string, 0)
		linkTexts := make([]string, 0)
		numberOfExternalScripts := 0
		numberOfInternalScripts := 0
		numberOfInternalCss := 0
		numberOfExternalCss := 0
		numberOfOutboundLinks := 0
		numberOfInboundLinks := 0

		e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
			words = append(words, strings.Fields(elem.Text)...)
		})

		e.ForEach("p", func(_ int, elem *colly.HTMLElement) {
			texts = append(texts, helpers.CleanUpString(elem.Text))
		})

		e.ForEach("img", func(_ int, elem *colly.HTMLElement) {
			imgSource := elem.Attr("src")
			url := document.Url.Scheme + "://" + document.Url.Host

			if imgSource != "" && helpers.IsAbsoluteUrl(imgSource) {
				images = append(images, imgSource)
			} else if imgSource != "" && !helpers.IsAbsoluteUrl(imgSource) {
				images = append(images, url+imgSource)
			}
		})

		e.ForEach("a", func(_ int, elem *colly.HTMLElement) {
			parsinghelper.CountNumberOfOutboundLinks(&numberOfOutboundLinks, elem)
			parsinghelper.CountNumberOfInboundLinks(&numberOfInboundLinks, elem)

			linkText := helpers.CleanUpString(elem.Text)
			if linkText == "" {
				return
			}
			linkTexts = append(linkTexts, linkText)
		})

		e.ForEach("script", func(_ int, elem *colly.HTMLElement) {
			parsinghelper.CountNumberOfExternalScripts(&numberOfExternalScripts, elem)
			parsinghelper.CountNumberOfInternalScripts(&numberOfInternalScripts, elem)
		})

		e.ForEach(`link[rel="stylesheet"]`, func(_ int, elem *colly.HTMLElement) {
			parsinghelper.CountNumberOfInternalCss(&numberOfInternalCss, elem)
			parsinghelper.CountNumberOfExternalCss(&numberOfExternalCss, elem)
		})

		document.NumberOfExternalCss = numberOfExternalCss
		document.NumberOfInternalCss = numberOfInternalCss
		document.NumberOfExternalScripts = numberOfExternalScripts
		document.NumberOfInternalScripts = numberOfInternalScripts
		document.NumberOfInboundLinks = numberOfInboundLinks
		document.NumberOfOutboundLinks = numberOfOutboundLinks

		document.LinkTexts = linkTexts
		document.Texts = texts
		document.Images = images
		document.WordsCount = parsinghelper.GetWordsCount(words)
	})

	c.OnHTML(`meta[property="og:type"]`, func(e *colly.HTMLElement) {
		document.OgTags.OgType = e.Attr("content")
	})

	c.OnHTML(`meta[name="google-site-verification"]`, func(e *colly.HTMLElement) {
		if e.Attr("content") != "" {
			document.HasGoogleSiteVerificationHtmlTag = true
		} else {
			document.HasGoogleSiteVerificationHtmlTag = false
		}
	})

	c.OnHTML(`meta[name="keywords"]`, func(e *colly.HTMLElement) {
		keywordsArray := strings.Split(e.Attr("content"), ",")
		document.Keywords = helpers.CleanUpStringsArray(keywordsArray)
	})

	c.OnHTML(`meta[name="viewport"]`, func(e *colly.HTMLElement) {
		if e.Attr("content") != "" {
			document.HasViewportHtmlTag = true
		} else {
			document.HasViewportHtmlTag = false
		}
	})

	c.OnHTML(`link[rel="icon"]`, func(e *colly.HTMLElement) {
		url := document.Url.Scheme + "://" + document.Url.Host
		document.FavIcon = url + e.Attr("href")
	})

	c.OnHTML(`link[rel="manifest"]`, func(e *colly.HTMLElement) {
		document.Manifest = e.Attr("href")
	})

	c.OnHTML(`meta[property="og:title"]`, func(e *colly.HTMLElement) {
		document.OgTags.OgTitle = e.Attr("content")
	})

	c.OnHTML(`meta[property="og:description"]`, func(e *colly.HTMLElement) {
		document.OgTags.OgDescription = helpers.CleanUpString(e.Attr("content"))
	})

	c.OnHTML(`meta[property="og:url"]`, func(e *colly.HTMLElement) {
		document.OgTags.OgUrl = e.Attr("content")
	})

	c.OnHTML(`meta[property="og:image"]`, func(e *colly.HTMLElement) {
		document.OgTags.OgImage = e.Attr("content")
	})

	c.OnHTML(`meta[property="og:site_name"]`, func(e *colly.HTMLElement) {
		document.OgTags.OgSiteName = e.Attr("content")
	})

	c.OnHTML(`html`, func(e *colly.HTMLElement) {
		document.Language = e.Attr("lang")
	})

	c.OnHTML(`meta[name="description"]`, func(e *colly.HTMLElement) {
		document.MetaTags.Description = helpers.CleanUpString(e.Attr("content"))
	})

	c.OnHTML(`meta[name="twitter:card"]`, func(e *colly.HTMLElement) {
		document.TwitterTags.Card = e.Attr("content")
	})

	c.OnHTML(`meta[name="twitter:title"]`, func(e *colly.HTMLElement) {
		document.TwitterTags.Title = helpers.CleanUpString(e.Attr("content"))
	})

	c.OnHTML(`meta[name="twitter:description"]`, func(e *colly.HTMLElement) {
		document.TwitterTags.Description = helpers.CleanUpString(e.Attr("content"))
	})

	c.OnHTML(`meta[name="twitter:domain"]`, func(e *colly.HTMLElement) {
		document.TwitterTags.Domain = e.Attr("content")
	})

	c.OnHTML(`meta[name="twitter:creator"]`, func(e *colly.HTMLElement) {
		document.TwitterTags.Creator = e.Attr("content")
	})

	c.OnHTML(`meta[name="twitter:image:src"]`, func(e *colly.HTMLElement) {
		document.TwitterTags.Image = e.Attr("content")
	})

	c.OnHTML(`link[rel="canonical"]`, func(e *colly.HTMLElement) {
		document.Canonical = e.Attr("href")
	})

	c.Visit(link)

	return *document
}
