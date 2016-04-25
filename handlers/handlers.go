package handlers

import (
	"net/http"
	"html/template"
	"path"
)

func Index(w http.ResponseWriter, r *http.Request){
	tmplpath := path.Join("templates", "index.html")
    tmpl, _ := template.ParseFiles(tmplpath)
    tmpl.Execute(w, nil)
	//fmt.Fprintf(w, "Batian 0.0.1")
}