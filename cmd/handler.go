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
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println("2: ", err)
		return
	}
}

func ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	textInput := r.FormValue("input_text")
	chooseStyle := r.FormValue("style")
	download := r.FormValue("download")

	asciiStyle, statusOfStyle := pkg.ChooseAsciiStyle(chooseStyle)
	if !statusOfStyle {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result, statusOfAscii := pkg.AsciiDrawer(textInput, asciiStyle)
	if !statusOfAscii {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if download == "Download" {
		w.Header().Set("Content-Disposition", "attachment; filename=result.txt")
		w.Header().Set("Conten-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(result)))
		_, err := w.Write([]byte(result))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}
