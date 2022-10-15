package main

import (
	"ascii-art-web/pkg"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type dataOfClient struct {
	Output string
	Input  string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		printErrors(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		printErrors(w, http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		printErrors(w, http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		printErrors(w, http.StatusInternalServerError)
		fmt.Println("2: ", err)
		return
	}
}

func ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		printErrors(w, http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		printErrors(w, http.StatusInternalServerError)
		return
	}

	textInput := r.FormValue("input_text")
	chooseStyle := r.FormValue("style")
	download := r.FormValue("download")

	asciiStyle, statusOfStyle := pkg.ChooseAsciiStyle(chooseStyle)
	if !statusOfStyle {
		printErrors(w, http.StatusInternalServerError)
		return
	}

	result, statusOfAscii := pkg.AsciiDrawer(textInput, asciiStyle)
	if !statusOfAscii {
		printErrors(w, http.StatusBadRequest)
		return
	}

	if download == "Download" {
		w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(result)))
		_, err := w.Write([]byte(result))
		if err != nil {
			printErrors(w, http.StatusInternalServerError)
			return
		}
	}
	data := &dataOfClient{
		Output: result,
		Input:  textInput,
	}
	err = ts.Execute(w, data)
	if err != nil {
		printErrors(w, http.StatusInternalServerError)
		return
	}
}
