package pudl

import "fmt"

var (
	SHOW_LOG = false
)

func log(entry string) {
	if SHOW_LOG {
		fmt.Println(entry)
	}
}

func EnableLogs() {
	SHOW_LOG = true
}
