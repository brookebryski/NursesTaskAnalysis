package routePack

import (
	"NurseTasks/dbconnection"
	"NurseTasks/models"
	"encoding/json"
	"net/http"
	"text/template"

	"golang.org/x/crypto/bcrypt"
)

type Sub struct {
	Username string
	Data     string
}

var tpl *template.Template

func Routing() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/nurselogin", getNurseFormHandler)
	http.HandleFunc("/nursetasks", loginauthFormHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/supervisorlogin", getSupFormHandler)
	http.HandleFunc("/suphome", getSupHomeHandler)
	http.HandleFunc("/nurseTasks", getNurseTasks)
	http.ListenAndServe(":3030", nil)
}

func loginauthFormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var user models.User
	username := r.FormValue("nurseUser")
	password := r.FormValue("nursePass")

	result := dbconnection.DB().Select("password", "name", "is_nurse", "id", "is_logged_in", "facility_id", "region_id", "is_supervisor", "is_logged_in").First(&user, "name = ?", username)

	if result.Error == nil {

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err == nil && user.Name == username {

			if user.IsNurse == "True" {
				user.IsLoggedIn = "True"
				tpl.ExecuteTemplate(w, "nurseTasks.html", user)

			} else {

				tpl.ExecuteTemplate(w, "nurseLogin.html", "You dont have access because you are not a Nurse")
				return
			}

		} else {
			tpl.ExecuteTemplate(w, "nurseLogin.html", "Your credentials dont match, please try again")
			return
		}
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var tasks []models.Task

	result := dbconnection.DB().Find(&tasks)
	if result.Error == nil {
		//jsonResp, err := json.Marshal(tasks)
		//if err != nil {
		//	log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		//}

		//w.Write(jsonResp)
		json.NewEncoder(w).Encode(tasks)
	}

}

func getNurseFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "nurseLogin.html", nil)
}

func getSupFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "supervisorLogin.html", nil)
}

func getSupHomeHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	var user models.User
	username := r.FormValue("nurseUser")
	password := r.FormValue("nursePass")

	result := dbconnection.DB().Select("password", "name", "is_nurse", "id", "is_logged_in", "facility_id", "region_id", "is_supervisor", "is_logged_in").First(&user, "name = ?", username)

	if result.Error == nil {

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

		if err == nil && user.Name == username {

			if (user.IsNurse == "True" && user.IsSupervisor == "True") || user.IsSupervisor == "True" {
				user.IsLoggedIn = "True"
				tpl.ExecuteTemplate(w, "supTasksView.html", user)

			} else {

				tpl.ExecuteTemplate(w, "supervisorLogin.html", "You dont have access because you are not a Supervisor")
				return
			}

		} else {
			tpl.ExecuteTemplate(w, "supervisorLogin.html", "Your credentials dont match, please try again")
			return
		}
	}

}

func getNurseTasks(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "nurseTasks.html", nil)
}
