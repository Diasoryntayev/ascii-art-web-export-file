package main

import (
	"html/template"
	"net/http"
)

type Response struct {
	NumberOfError int
	Message       string
}

func printErrors(w http.ResponseWriter, code int) {
	ts, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := Response{
		NumberOfError: code,
		Message:       http.StatusText(code),
	}
	w.WriteHeader(resp.NumberOfError)
	err = ts.Execute(w, resp)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
