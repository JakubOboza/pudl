package pudl

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

const (
	TIMEOUT = 2
)

var (
	ErrOutOfCacheBounds = errors.New("index provided by user is out of usable bounds")
)

func Show(puldId int) error {

	// load config
	config, err := LoadConfigFromFile()

	if err != nil {
		log(err.Error())
		return err
	}

	dbIndex := puldId - 1

	if dbIndex < 0 || dbIndex > len(config.Data) {
		log("Index out of bounds")
		return ErrOutOfCacheBounds
	}

	pudl := config.Data[dbIndex]

	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)

	defer cancel()

	cmd := exec.CommandContext(ctx, openForUserOs(), pudl.Url)

	if ctx.Err() == context.DeadlineExceeded {
		log("open command timeout exceeded")
		return ctx.Err()
	}

	log(fmt.Sprintf("Visiting %s", pudl.Url))

	err = cmd.Run()

	if err != nil {
		return err
	}

	return nil
}

func openForUserOs() string {
	defaultOption := "open"

	switch runtime.GOOS {
	case "darwin":
		return "open"
	case "linux":
		return "xdg-open"
	case "freebsd":
		return "xdg-open" //Not checked yet
	case "windows":
		return "explorer"
	default:
		return defaultOption
	}
}
