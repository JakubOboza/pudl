package pudl

import (
	"errors"
	"fmt"

	"github.com/gocolly/colly/v2"
)

const (
	SELECTOR_BASE       = "[data-st-area=main-stream-%d]"
	MEDIA_STREAMS_COUNT = 21
)

var (
	ErrNothingFound = errors.New("Did not found anything on the page. Maybe structure has changed! Please update pudl")
)

func ScrapePudelekFrontPage() (*[]Pudl, error) {

	pudls := []Pudl{}

	c := colly.NewCollector()

	for i := 1; i < MEDIA_STREAMS_COUNT; i++ {
		selector := fmt.Sprintf(SELECTOR_BASE, i)

		c.OnHTML(selector, func(e *colly.HTMLElement) {
			title := e.ChildText("a")
			hrefPath := e.ChildAttr("a", "href")

			newPudl := Pudl{
				Title: title,
				Url:   pudelekUrl(hrefPath),
			}
			pudls = append(pudls, newPudl)
		})

	}

	err := c.Visit(PUDELEK_BASE_URL)

	if err != nil {
		return nil, err
	}

	if len(pudls) == 0 {
		return nil, ErrNothingFound
	}

	return &pudls, nil
}
