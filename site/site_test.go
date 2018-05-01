package site

import "testing"

func TestCorrectlyDomainExtraction(t *testing.T) {
	assertEqual(t, "thisisdeadpool.com", getSiteHostname("http://thisisdeadpool.com/cable-deadpool-001-2004/"))
	assertEqual(t, "thisisdeadpool.com", getSiteHostname("https://thisisdeadpool.com/cable-deadpool-001-2004/"))
	assertEqual(t, "thisisdeadpool.com", getSiteHostname("//thisisdeadpool.com/cable-deadpool-001-2004/"))
	assertEqual(t, "thisisdeadpool.com", getSiteHostname("http://thisisdeadpool.com/cable-deadpool-001-2004/?query=params"))
}
