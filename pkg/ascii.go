package pkg

import (
	"log"
	"os"
	"strings"
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
	asciiArt := CreateMapWithAsciiArt(banner)
	res := OutputAscii(inputText, asciiArt)
	return res, status
}

func isOnlyAsciiSymbol(text string) ([]string, bool) {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	arrText := strings.Split(text, "\n")
	for _, v := range text {
		if (v < ' ' || v > '~') && v != '\n' {
			return nil, false
		}
	}
	return arrText, true
}
