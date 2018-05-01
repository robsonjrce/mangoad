package site

import (
	"fmt"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

var doc *goquery.Document

func assertEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected string to be the same. Got '%s' instead '%s'", actual, expected)
	}
}

func assertEqualPages(t *testing.T, expected []string, actual []string) {
	if len(expected) != len(actual) {
		t.Errorf("Expected same len. Got '%d' instead '%d'", len(actual), len(expected))
		return
	}
	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("Expected values to be the same. Got '%d' on index %d instead '%d'", len(actual), i, len(expected))
		}
	}
}

func loadDoc(page string) *goquery.Document {
	var f *os.File
	var e error

	if f, e = os.Open(fmt.Sprintf("../testdata/%s", page)); e != nil {
		panic(e.Error())
	}
	defer f.Close()

	var node *html.Node
	if node, e = html.Parse(f); e != nil {
		panic(e.Error())
	}
	return goquery.NewDocumentFromNode(node)
}

func Doc() *goquery.Document {
	if doc == nil {
		doc = loadDoc("thisisdeadpool-volume.html")
	}
	return doc
}

func TestValidHtmlDocument(t *testing.T) {
	thisisdeadpool := NewThisIsDeadpool("thisisdeadpool.com")

	thisisdeadpool.UpdateDocument(Doc())

	assertEqual(t, "Cable & Deadpool 001 (2004)", thisisdeadpool.GetVolumeTitle())

	expectedPages := []string{
		"http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-000.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-001.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-002.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-003.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-004.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-005.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-006.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-007.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-008.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-009.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-010.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-011.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-012.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-013.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-014.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-015.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-016.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-017.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-018.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-019.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-020.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-021.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-022.jpg",
	}

	assertEqualPages(t, expectedPages, thisisdeadpool.GetVolumePages())
}
