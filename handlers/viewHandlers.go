package handlers

import (
	"os"
	"net/http"
	"path"
	"html/template"
)

func AppDashboard(w http.ResponseWriter, r *http.Request){
	cwd, _ := os.Getwd()
	tmpl := template.Must(
		template.ParseFiles(path.Join(cwd, "templates/base.html"), path.Join(cwd, "templates/dashboard.html")))
	tmpl.Execute(w, nil)
}