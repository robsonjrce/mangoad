package site

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func assertEqual(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("Expected string to be the same. Got '%s' instead '%s'", actual, expected)
	}
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Expected 'nil'. Got type '%s'", reflect.TypeOf(actual).Name())
	}
}

func assertNotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Errorf("Expected error. Got 'nil'")
	}
}

func assertEqualArray(t *testing.T, expected []string, actual []string) {
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

func DocVolume() *goquery.Document {
	var doc *goquery.Document

	if doc == nil {
		doc = loadDoc("thisisdeadpool-volume.html")
	}
	return doc
}

func DocVolumes() *goquery.Document {
	var doc *goquery.Document

	if doc == nil {
		doc = loadDoc("thisisdeadpool-volumes.html")
	}
	return doc
}

func TestValidThisIsDeadpoolCreation(t *testing.T) {
	siteScraper, _ := NewSiteScraper("http://thisisdeadpool.com/cable-deadpool-001-2004/")

	if reflect.TypeOf(siteScraper).Kind() != reflect.TypeOf(&ThisIsDeadpool{}).Kind() {
		t.Error("error")
	}
}

func TestValidVolumeTitle(t *testing.T) {
	thisisdeadpool := NewThisIsDeadpool("thisisdeadpool.com")
	thisisdeadpool.UpdateDocument(DocVolume())

	assertEqual(t, "Cable & Deadpool 001 (2004)", thisisdeadpool.GetVolumeTitle())
}

func TestCorrectVolumePages(t *testing.T) {
	thisisdeadpool := NewThisIsDeadpool("thisisdeadpool.com")
	thisisdeadpool.UpdateDocument(DocVolume())

	expectedPages := []string{
		"http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-000.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-001.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-002.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-003.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-004.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-005.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-006.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-007.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-008.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-009.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-010.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-011.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-012.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-013.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-014.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-015.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-016.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-017.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-018.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-019.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-020.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-021.jpg", "http://thisisdeadpool.com/wp-content/uploads/2017/12/Cable-Deadpool-001-022.jpg",
	}

	assertEqualArray(t, expectedPages, thisisdeadpool.GetVolumePages())
}

func TestCorrectVolumesFromList(t *testing.T) {
	thisisdeadpool := NewThisIsDeadpool("thisisdeadpool.com")
	thisisdeadpool.UpdateDocument(DocVolumes())

	expectedVolumes := []string{
		"http://thisisdeadpool.com/cable-deadpool-001-2004/", "http://thisisdeadpool.com/cable-deadpool-002-2004/", "http://thisisdeadpool.com/cable-deadpool-003-2004/", "http://thisisdeadpool.com/cable-deadpool-004-2004/", "http://thisisdeadpool.com/cable-deadpool-005-2004/", "http://thisisdeadpool.com/cable-deadpool-006-2004/", "http://thisisdeadpool.com/cable-deadpool-007-2004/", "http://thisisdeadpool.com/cable-deadpool-008-2004/", "http://thisisdeadpool.com/cable-deadpool-009-2005/", "http://thisisdeadpool.com/cable-deadpool-010-2005/", "http://thisisdeadpool.com/cable-deadpool-011-2005/", "http://thisisdeadpool.com/cable-deadpool-012-2005/", "http://thisisdeadpool.com/cable-deadpool-013-2005/", "http://thisisdeadpool.com/cable-deadpool-014-2005/", "http://thisisdeadpool.com/cable-deadpool-015-2005/", "http://thisisdeadpool.com/cable-deadpool-016-2005/", "http://thisisdeadpool.com/cable-deadpool-017-2005/", "http://thisisdeadpool.com/cable-deadpool-018-2005/", "http://thisisdeadpool.com/cable-deadpool-019-2005/", "http://thisisdeadpool.com/cable-deadpool-020-2005/", "http://thisisdeadpool.com/cable-deadpool-021-2005/", "http://thisisdeadpool.com/cable-deadpool-022-2006/", "http://thisisdeadpool.com/cable-deadpool-023-2006/", "http://thisisdeadpool.com/cable-deadpool-024-2006/", "http://thisisdeadpool.com/cable-deadpool-025-2006/", "http://thisisdeadpool.com/cable-deadpool-026-2006/", "http://thisisdeadpool.com/cable-deadpool-027-2006/", "http://thisisdeadpool.com/cable-deadpool-028-2006/", "http://thisisdeadpool.com/cable-deadpool-029-2006/", "http://thisisdeadpool.com/cable-deadpool-030-2006/", "http://thisisdeadpool.com/cable-deadpool-031-2006/", "http://thisisdeadpool.com/cable-deadpool-032-2006/", "http://thisisdeadpool.com/cable-deadpool-033-2006/", "http://thisisdeadpool.com/cable-deadpool-034-2007/", "http://thisisdeadpool.com/cable-deadpool-035-2007/", "http://thisisdeadpool.com/cable-deadpool-036-2007/", "http://thisisdeadpool.com/cable-deadpool-037-2007/", "http://thisisdeadpool.com/cable-deadpool-038-2007/", "http://thisisdeadpool.com/cable-deadpool-039-2007/", "http://thisisdeadpool.com/cable-deadpool-040-2007/", "http://thisisdeadpool.com/cable-deadpool-041-2007/", "http://thisisdeadpool.com/cable-deadpool-042-2007/", "http://thisisdeadpool.com/cable-deadpool-043-2007/", "http://thisisdeadpool.com/cable-deadpool-044-2007/", "http://thisisdeadpool.com/cable-deadpool-045-2007/", "http://thisisdeadpool.com/cable-deadpool-046-2007/", "http://thisisdeadpool.com/cable-deadpool-047-2008/", "http://thisisdeadpool.com/cable-deadpool-048-2008/", "http://thisisdeadpool.com/cable-deadpool-049-2008/", "http://thisisdeadpool.com/cable-deadpool-050-2008/",
	}

	assertEqualArray(t, expectedVolumes, thisisdeadpool.GetVolumes())
}
