package pudl

import (
	"fmt"
	"strings"
)

func Display(pudlsToShow *[]Pudl, oldPudls *[]Pudl, maxCount int) {
	fmt.Println("⭐⭐⭐ 😹 Lista tematow z pudla 🐩 😹 ⭐⭐⭐")
	for index, pudl := range *pudlsToShow {
		if maxCount > 0 && maxCount == index {
			break
		}
		existsInCache := existsCachedByTitle(pudl.Title, oldPudls)
		fmt.Printf("%d:%s %s\n", index+1, markIfNew(existsInCache), cleanupTitle(pudl.Title))
	}
}

func existsCachedByTitle(title string, oldPudls *[]Pudl) bool {
	if oldPudls == nil {
		return false
	}
	found := false

	for _, pudl := range *oldPudls {
		if pudl.Title == title {
			found = true
		}
	}

	return found
}

func markIfNew(exists bool) string {
	if exists {
		return ""
	} else {
		return "⭐"
	}
}

func cleanupTitle(title string) string {
	return strings.Replace(title, "Zobacz więcej", "", -1)
}
