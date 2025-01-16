package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

type sheetdetails struct {
	geometry string
	user     string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := sheetdetails{
			geometry: r.FormValue("geometry"),
			user:     r.FormValue("user"),
		}

		// do something with details
		//processing for the sheet geometry goes here. Can call Python/Matlab bindings in another file
		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}
