package site

import (
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Site define common information to supported sites modules
type Site interface {
	GetURL() string
	GetPageTitle() string
	GetPages() []string
	UpdateDocument(doc *goquery.Document)
}

func getSiteContent(address string) *goquery.Document {
	// Request the HTML page.
	res, err := http.Get(address)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

func getSiteHostname(address string) string {
	u, err := url.Parse(address)
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(u.Hostname(), ".")
	domain := parts[len(parts)-2] + "." + parts[len(parts)-1]

	return domain
}

// NewSiteScraper will factory instantiate the right site
func NewSiteScraper(url string) {
	domain := getSiteHostname(url)
	if domain != "" {
		println(domain)
	}

	for _, module := range modules {
		if module.ID == domain {
			site := module.Fn(url)
			println(site.GetURL())
			println(reflect.TypeOf(site).Name())
		}
	}
}
