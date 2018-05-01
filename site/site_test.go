package site

import "testing"

func TestCorrectlyDomainExtraction(t *testing.T) {
	assertEqual(t, "thisisdeadpool.com", getSiteHostname("http://thisisdeadpool.com/cable-deadpool-001-2004/"))
	assertEqual(t, "thisisdeadpool.com", getSiteHostname("https://thisisdeadpool.com/cable-deadpool-001-2004/"))
	assertEqual(t, "thisisdeadpool.com", getSiteHostname("//thisisdeadpool.com/cable-deadpool-001-2004/"))
	assertEqual(t, "thisisdeadpool.com", getSiteHostname("http://thisisdeadpool.com/cable-deadpool-001-2004/?query=params"))
}

func TestDomainNotSupported(t *testing.T) {
	siteScraper, err := NewSiteScraper("http://thissiteisnotsupported.com")

	assertNil(t, siteScraper)
	assertNotNil(t, err)
	assertEqual(t, "site is not supported", err.Error())
}
