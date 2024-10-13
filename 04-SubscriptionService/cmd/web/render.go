package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var pathToTemplates = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]interface{}
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
}

func (app *Config) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTemplates),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplates),
	}

	var templateSlice []string

	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplates, t))
	templateSlice = append(templateSlice, partials...)

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, app.AddDefaultData(td, r)); err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Error = app.Sessions.PopString(r.Context(), "error")
	td.Flash = app.Sessions.PopString(r.Context(), "flash")
	td.Warning = app.Sessions.PopString(r.Context(), "warning")
	if app.isAuthenticated(r) {
		td.Authenticated = true
		// TODO - get more user information
	}
	td.Now = time.Now()
	return td
}

func (app *Config) isAuthenticated(r *http.Request) bool {
	return app.Sessions.Exists(r.Context(), "userID")
}
