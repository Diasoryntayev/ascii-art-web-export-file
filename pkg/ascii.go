package pkg

import (
	"log"
	"os"
)

func ChooseAsciiStyle(s string) (string, bool) {
	var r string
	switch s {
	case "standard":
		r = "pkg/banners/standard.txt"
	case "shadow":
		r = "pkg/banners/shadow.txt"
	case "thinkertoy":
		r = "pkg/banners/thinkertoy.txt"
	}
	_, err := os.Stat(r)
	if err != nil {
		log.Println(err)
		return "", false
	}
	return r, true
}

func AsciiDrawer(input, banner string) (string, bool) {
	inputText, status := isOnlyAsciiSymbol(input)
	if !status {
		return "", status
	}
}
