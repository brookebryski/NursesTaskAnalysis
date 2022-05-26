package routePack

import (
	"net/http"
	"text/template"
)

type Sub struct {
	Username string
	Data     string
}

var tpl *template.Template

func Routing() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/nurseLogin", getNurseFormHandler)
	http.HandleFunc("/supLogin", getSupFormHandler)
	http.HandleFunc("/supHome", getSupHomeHandler)
	http.HandleFunc("/nurseTasks", getNurseTasks)
	http.ListenAndServe(":3030", nil)
}

func getNurseFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "nurseLogin.html", nil)
}

func getSupFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "supLogin.html", nil)
}

func getSupHomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "supTasksView.html", nil)
}

func getNurseTasks(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "nurseTasks.html", nil)
}
