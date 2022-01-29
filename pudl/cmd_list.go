package pudl

import (
	"fmt"
	"time"
)

func List() error {

	// load config
	config, err := LoadConfigFromFile()

	if err != nil {
		log(err.Error())
		return err
	}

	log(fmt.Sprintf("Last checked %v", time.Now().Sub(config.LastFetchTime)))

	// check cache and time
	if len(config.Data) > 1 && time.Now().Sub(config.LastFetchTime) < time.Second*REFRESH_THROTTLE {
		// Display Cached values
		log("Using cache")
		Display(&config.Data, &config.Data, -1)
	} else {
		// fetch
		pudls, err := ScrapePudelekFrontPage()

		if err != nil {
			return err
		}

		config.LastFetchTime = time.Now()

		if len(*pudls) > 0 {
			Display(pudls, &config.Data, -1)
		} else {
			return ErrNothingFound
		}
		config.Data = *pudls
	}

	return config.SaveToFile()
}
