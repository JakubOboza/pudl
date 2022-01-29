package pudl

import (
	"testing"
)

/*
This is an integration test to detect
if structure didn't change and to prove scrapper works
*/

func TestScrapper(t *testing.T) {

	pudls, err := ScrapePudelekFrontPage()

	if err != nil {
		t.Fatal(err)
	}

	if len(*pudls) < 2 {
		t.Errorf("Expected at least 2 entries but got '%d'", len(*pudls))
	}

}
