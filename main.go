package main

import (
	"flag"
	"fmt"
	"net/url"
	"site"
	"strings"
)

// UrlFlag is a custom flag type to import url
type UrlFlag struct {
	Urls []*url.URL
}

func (arr *UrlFlag) String() string {
	return fmt.Sprint(arr.Urls)
}

func (arr *UrlFlag) Set(value string) error {
	urls := strings.Split(value, ",")
	for _, item := range urls {
		if parsedUrl, err := url.Parse(item); err != nil {
			return err
		} else {
			arr.Urls = append(arr.Urls, parsedUrl)
		}
	}
	return nil
}

var arg UrlFlag
var urlType string

func main() {
	flag.Var(&arg, "url", "URL comma-separated list")
	flag.StringVar(&urlType, "t", "v", "type of the url being tracked [s serie, l list, v volume]")
	flag.Parse()

	for _, item := range arg.Urls {
		fmt.Printf("scheme: %s url: %s path: %s\n", item.Scheme, item.Host, item.Path)
	}

	site.NewSiteScraperFromConfig("http://thisisdeadpool.com/cable-deadpool-001-2004/", "k")

	// _, err := site.NewSiteScraper("http://thisisdeadpool.com/cable-deadpool-001-2004/")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// config.GetConfigContent()

	// for _, page := range scraper.GetVolumePages() {
	// 	println(page)
	// }
}
