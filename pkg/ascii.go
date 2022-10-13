package pkg

import (
	"log"
	"os"
)

func chooseAsciiStyle(s string) (string, bool) {
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
