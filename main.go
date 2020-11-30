package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", firstHandle)
	http.ListenAndServe(":8080", nil)
}

func firstHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		http.ServeFile(w, r, "index.html")
	} else if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		form := r.FormValue("form")
		text := r.FormValue("textInput")
		fmt.Fprintf(w, "form : %s\n", form)
		fmt.Fprintf(w, "text : %s\n", text)
	} else {
		fmt.Fprintf(w, "Sorry")
	}
}
