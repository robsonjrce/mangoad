package site

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ThisIsDeadpool supports the site of same name
type ThisIsDeadpool struct {
	URL string
	Doc *goquery.Document
}

// GetURL will return url
func (s *ThisIsDeadpool) GetURL() string {
	return s.URL
}

// UpdateDocument will update document data
func (s *ThisIsDeadpool) UpdateDocument(doc *goquery.Document) {
	s.Doc = doc
}

// GetPageTitle will return
func (s *ThisIsDeadpool) GetPageTitle() (title string) {
	s.Doc.Find(".site .site-content .content-area .site-main .post .entry-header").Each(func(i int, s *goquery.Selection) {
		title = s.Find("h1").Text()
		title = strings.TrimSpace(title)
	})
	return
}

// GetPages will return
func (s *ThisIsDeadpool) GetPages() (pages []string) {
	// Find the review items
	s.Doc.Find(".site .site-content .content-area .site-main .post .entry-content p img").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		img, exist := s.Attr("src")
		if exist {
			pages = append(pages, img)
		}
	})
	return
}

// NewThisIsDeadpool will instantiate a new module
func NewThisIsDeadpool(url string) Site {
	return &ThisIsDeadpool{URL: url}
}

func init() {
	RegisterModule("thisisdeadpool.com", NewThisIsDeadpool)
}
