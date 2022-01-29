package pudl

import (
	"testing"
)

func TestEnableLogs(t *testing.T) {
	SHOW_LOG = false
	EnableLogs()
	if SHOW_LOG != true {
		t.Errorf("Expected EnableLogs to set sHOW_LOG to true but got '%v'", SHOW_LOG)
	}
}
