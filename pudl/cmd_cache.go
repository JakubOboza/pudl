package pudl

import (
	"errors"
	"fmt"
)

var (
	ErrNothingCached = errors.New("Nothing is cached yet")
)

func Cache() error {

	// load config
	config, err := LoadConfigFromFile()

	if err != nil {
		log(err.Error())
		return err
	}

	if len(config.Data) < 1 {
		return ErrNothingCached
	}

	fmt.Printf("Last fetch %v\n", config.LastFetchTime)
	Display(&config.Data, &config.Data, -1)

	return nil
}
