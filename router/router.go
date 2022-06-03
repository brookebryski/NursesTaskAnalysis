package router

import (
	"NursesTaskAnalysis/dbconnection"
	"NursesTaskAnalysis/models"
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Sub struct {
	Username string
	Data     string
}

var tpl *template.Template

func Routing() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/nurselogin", getNurseFormHandler)                                      //login page for nurses.
	http.HandleFunc("/nursetasks", loginauthFormHandler)                                     //nurses task/home page. Validation occurs here
	http.HandleFunc("/tasks", tasksHandler)                                                  //nurse page, populates the auto_select dropdown for submit a task form
	http.HandleFunc("/supervisorlogin", getSupFormHandler)                                   //login page for supervisor
	http.HandleFunc("/suphome", getSupHomeHandler)                                           //home page for supervisor, validation occurs here
	http.HandleFunc("/", postTask)                                                           //for nurses page, creates a new task assigned to a department
	http.HandleFunc("/getenteredtasks", getEnteredTasks)                                     //for nurse page, gets the task table for completed tasks
	http.HandleFunc("/getfacilitiesdropdowndatabyregion", getFacilitiesDropdownDataByRegion) //for supervisors, populates dropdown for facilities based on selected region dropdown
	http.HandleFunc("/getfacilitiestabledatabyregion", getFacilitiesTableDataByRegion)       //cut from final view
	http.HandleFunc("/createnewtask", createNewTasks)                                        //for nurses, takes a modal of whatever the nurse enters and creates a new task that can be submitted.
	http.HandleFunc("/getregionandfacility", getregionandfacilityHandler)                    //grabs the region Name of the user and the facility Name of the user and displays it on information
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
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var tasks []models.Task
	id := r.URL.Query().Get("id")
	result := dbconnection.DB().Find(&tasks, "department_id = ?", id)
	if result.Error == nil {
		jsonResp, err := json.Marshal(tasks)
		if err != nil {
			panic(err)
		}

		w.Write(jsonResp)
		//json.NewEncoder(w).Encode(tasks)
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

func postTask(w http.ResponseWriter, r *http.Request) {
	var t models.TaskEntered
	w.WriteHeader(http.StatusCreated)

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := dbconnection.DB().Create(&t)
	if result.Error != nil {
		panic(err)
	}

}

func getEnteredTasks(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	type Task struct {
		CreatedAt    time.Time
		TaskID       uint   `gorm:"not null"`
		UserID       uint   `gorm:"not null"`
		ID           uint   `gorm:"not null"`
		Name         string `gorm:"not null"`
		DepartmentID uint   `gorm:"not null"`
	}
	var tasks []Task
	id := r.URL.Query().Get("id")
	result := dbconnection.DB().Select("task_entereds.created_at,task_entereds.task_id,task_entereds.user_id,tasks.*").Joins("JOIN task_entereds on task_entereds.task_id=tasks.id").Find(&tasks, "user_id = ?", id)
	if result.Error == nil {
		jsonResp, err := json.Marshal(tasks)
		if err != nil {
			panic(err)
		}
		w.Write(jsonResp)

	}

}

func getFacilitiesDropdownDataByRegion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	type Facility struct {
		ID       uint   `gorm:"not null"`
		RegionID uint   `gorm:"not null"`
		Name     string `gorm:"not null"`
	}
	var facilities []Facility
	id := r.URL.Query().Get("id")
	result := dbconnection.DB().Select("facilities.id,facilities.region_id,facilities.name").Joins("JOIN regions on facilities.region_id=regions.id").Find(&facilities, "region_id = ?", id)
	if result.Error == nil {
		jsonResp, err := json.Marshal(facilities)
		if err != nil {
			panic(err)
		}

		w.Write(jsonResp)

	}

}

func getFacilitiesTableDataByRegion(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	type Facility struct {
		ID       uint   `gorm:"not null"`
		RegionID uint   `gorm:"not null"`
		Name     string `gorm:"not null"`
	}
	var facilities []Facility
	id := r.URL.Query().Get("id")
	result := dbconnection.DB().Select("facilities.id,facilities.region_id,facilities.name").Joins("JOIN regions on facilities.region_id=regions.id").Find(&facilities, "region_id = ?", id)
	if result.Error == nil {
		jsonResp, err := json.Marshal(facilities)
		if err != nil {
			panic(err)
		}

		w.Write(jsonResp)

	}
}

func createNewTasks(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	w.WriteHeader(http.StatusCreated)
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := dbconnection.DB().Create(&t)
	if result.Error != nil {
		panic(err)
	}

}

func getregionandfacilityHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	type Region struct {
		Name string `gorm:"not null"`
	}
	type Facility struct {
		Name string `gorm:"not null"`
	}
	type RegionFacility struct {
		RName string `gorm:"not null"`
		FName string `gorm:"not null"`
	}
	var region Region
	var facility Facility
	var regionFacility RegionFacility
	fid := r.URL.Query().Get("fid")
	rid := r.URL.Query().Get("rid")

	result1 := dbconnection.DB().Select("regions.name").Find(&region, "regions.id = ?", rid)
	if result1.Error == nil {

		regionFacility.RName = region.Name
	}
	result2 := dbconnection.DB().Select("facilities.name").Find(&facility, "facilities.id = ?", fid)
	if result2.Error == nil {

		regionFacility.FName = facility.Name

	}
	jsonResp, err := json.Marshal(regionFacility)
	if err != nil {
		panic(err)
	}
	w.Write(jsonResp)

}
