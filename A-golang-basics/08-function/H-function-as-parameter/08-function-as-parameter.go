package main

import (
	"fmt"
	"strings"
)

// declaring function type
type Filter func(string) string

func printComment(comment string, filter Filter) {
	filteredComment := filter(comment)
	fmt.Println(filteredComment)
}

func buzzWordFilter(comment string) string {
	filtered := strings.ToLower(comment)
	buzz := [...]string{"anjing", "tai", "nigga", "tolol", "cyna", "nigger", "batjingan"}

	for _, e := range buzz {
		filtered = strings.ReplaceAll(filtered, e, "..!")
	}

	return filtered
}

func main() {
	printComment("Dasar Cyna tolol xialan", buzzWordFilter)
	printComment("Dasar ni9ger tolol batjingan", buzzWordFilter)
}
