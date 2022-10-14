package pkg

import (
	"bufio"
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

func CreateMapWithAsciiArt(AsciiFont string) map[rune][]string {
	var s string
	var counter uint8
	data, _ := os.Open(AsciiFont)
	artText := bufio.NewScanner(data)
	m := map[rune][]string{}
	i := ' '
	for artText.Scan() {
		s = artText.Text()
		if s != "" {
			m[i] = append(m[i], s)
			counter++
		}
		if counter == 8 {
			counter = 0
			i++
		}
	}
	m['\n'] = []string{"", "", "", "", "", "", "", ""}
	return m
}

func OutputAscii(clientText []string, ascii map[rune][]string) string {
	var result string
	for j := 0; j < len(clientText); j++ {
		if clientText[j] == "" {
			result += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for _, v := range clientText[j] {
				result += (ascii[(v)][i])
			}
			result += "\n"
		}
	}
	return result
}
