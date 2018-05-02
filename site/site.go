package site

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var (
	httpc = &http.Client{
		Timeout: time.Millisecond * 1000,
	}
)

// Site define common information to supported sites modules
type Site interface {
	GetURL() string
	GetVolumeTitle() string
	GetVolumePages() []string
	GetVolumes() []string
	UpdateDocument(doc *goquery.Document)
}

func getSiteContent(address string) *goquery.Document {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/45.0.2454.101 Chrome/45.0.2454.101 Safari/537.36")

	// Request the HTML page.
	res, err := httpc.Do(req)
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
func NewSiteScraper(url string) (site Site, err error) {
	domain := getSiteHostname(url)

	for _, module := range modules {
		if module.ID == domain {
			site = module.Fn(url)
			break
		}
	}

	if site == nil {
		err = errors.New("site is not supported")
	}
	return
}

// NewSiteScraperFromConfig is capable to factory the right scraper with minor customization
func NewSiteScraperFromConfig(url string, urlType string) (site Site, err error) {
	site, err = NewSiteScraper(url)
	if err != nil {
		return
	}

	// reflection all the way up so we can update default values
	ptyp := reflect.TypeOf(site)
	pval := reflect.ValueOf(site)

	var typ reflect.Type
	var val reflect.Value
	if ptyp.Kind() == reflect.Ptr {
		// argument is a pointer, dereferencing.
		typ = ptyp.Elem()
		val = pval.Elem()
	} else {
		typ = ptyp
		val = pval
	}
	if typ.Kind() != reflect.Struct {
		site = nil
		err = errors.New("it is not a struct")
		return
	}
	if val.CanSet() {
		log.Printf("We can set values")
	} else {
		site = nil
		err = errors.New("we cannot set values")
		return
	}

	for i := 0; i < typ.NumField(); i++ {
		// informations for type
		sfld := typ.Field(i)
		nfld := sfld.Name
		tfld := sfld.Type
		kfld := tfld.Kind()

		// informations for value
		vfld := val.Field(i)
		if vfld.CanSet() {
			switch nfld {
			case "Type":
				vfld.SetString(urlType)
			}
		}
		log.Printf("struct field '%d': name '%s' type '%s' kind '%s' value '%v'\n",
			i,
			sfld.Name,
			tfld,
			kfld,
			vfld)
	}
	return
}
