package scrapeutils

type domain struct {
	Link string
}

func GetDomains() (domains []domain) {

	var dms = []domain{
		{
			Link: "https://www.google.com",
		},
	}

	return dms
}
