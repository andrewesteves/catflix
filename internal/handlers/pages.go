package handlers

import (
	"net/http"
	"text/template"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./web/views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
