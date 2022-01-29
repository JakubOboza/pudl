package pudl

import (
	"net/url"
	"path"
	"regexp"
)

const (
	PUDELEK_BASE_URL = "https://pudelek.pl"
	PUDL_CONFIG_FILE = ".pudl.yaml"
)

type Pudl struct {
	Title string
	Url   string
}

func pudelekUrl(pudelPath string) string {
	if match, _ := regexp.MatchString("^http[s]?://", pudelPath); match {
		return pudelPath
	} else {
		u, _ := url.Parse(PUDELEK_BASE_URL)
		u.Path = path.Join(u.Path, pudelPath)
		return u.String()
	}
}
