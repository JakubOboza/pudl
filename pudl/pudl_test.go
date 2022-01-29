package pudl

import (
	"testing"
)

func TestPudelekUrl(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "",
			expected: "https://pudelek.pl",
		},
		{
			input:    "/a/b/c/d",
			expected: "https://pudelek.pl/a/b/c/d",
		},
		{
			input:    "a/b/c/d.lol",
			expected: "https://pudelek.pl/a/b/c/d.lol",
		},
		{
			input:    "http://wp.onet.interia.pl/lulzmuch",
			expected: "http://wp.onet.interia.pl/lulzmuch",
		},
	}

	for _, testCase := range testCases {

		got := pudelekUrl(testCase.input)

		if got != testCase.expected {
			t.Errorf("Expected '%s' but got '%s'", testCase.expected, got)
		}

	}

}
