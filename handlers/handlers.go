package handlers

import (
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request){
    tmpl, _ := template.ParseFiles("templates/base.html", "templates/index.html")
    tmpl.Execute(w, nil)
}

func SignIn(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case "GET":
			tmpl, _ := template.ParseFiles("templates/base.html", "templates/signin.html")
			tmpl.Execute(w, nil)
		case "POST":
			//r.ParseForm()
			//r.Form.Get("username")
			tmpl, _ := template.ParseFiles("templates/base.html", "templates/signin.html")
			tmpl.Execute(w, nil)
	}
	
}