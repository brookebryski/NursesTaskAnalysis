package router

import (
	"NursesTaskAnalysis/dbconnection"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type TaskEntered struct {
	TaskID int
	UserID int
}

type Task struct {
	Name string
	DepartmentID int
}

type User struct {
	ID           int   
	Name         string 
	Password     string 
	IsNurse      string 
	IsSupervisor string 
	RegionID     int   
	FacilityID   int   
	IsLoggedIn   string 
}
func TestMain(m *testing.M) {
	setupTest()
	code := m.Run()
	// teardown
	os.Exit(code)
}

func setupTest() {
	driver := dbconnection.Postgres
	dbconnection.New(driver)
}

var nurse = User{ 
	Name:"Helen",
    Password:"Helen123",
	IsNurse:"True",
	IsSupervisor: "False",
	RegionID: 1,
	FacilityID: 2,
	IsLoggedIn: "False",
}
func TestLoginauthFormHandler(t *testing.T) {
	subTests := []struct {
		name string
		body  User     // whatever obj you need
		expectedCode int
	}{
		{
			name: "SuccessfulCreateRequest",
			body: nurse,
			expectedCode: 200,
		},
	}
	
	for _, subTest := range subTests {
		t.Run(subTest.name, func(t *testing.T) {
			bodyJson, _ := json.Marshal(subTest.body)
			body := bytes.NewReader(bodyJson)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/nursetasks", body) // body if exists
			loginauthFormHandler(w, req)
			// res := w.Result() // if you need data
			if subTest.expectedCode != w.Code {
				t.Errorf("expected status code %v, got status code %v", subTest.expectedCode, w.Code)
			}
		})
	}
}

var supervisor = User{ 
	Name:"Patricia",
    Password:"Patricia123",
	IsNurse:"True",
	IsSupervisor: "True",
	RegionID: 1,
	FacilityID: 1,
	IsLoggedIn: "False",
}
func TestGetSupHomeHandler(t *testing.T) {
	subTests := []struct {
		name string
		body  User     // whatever obj you need
		expectedCode int
	}{
		{
			name: "SuccessfulCreateRequest",
			body: supervisor,
			expectedCode: 200,
		},
	}
	
	for _, subTest := range subTests {
		t.Run(subTest.name, func(t *testing.T) {
			bodyJson, _ := json.Marshal(subTest.body)
			body := bytes.NewReader(bodyJson)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/suphome", body) // body if exists
			getSupHomeHandler(w, req)
			// res := w.Result() // if you need data
			if subTest.expectedCode != w.Code {
				t.Errorf("expected status code %v, got status code %v", subTest.expectedCode, w.Code)
			}
		})
	}
}
func TestTasksHandler(t *testing.T) {
	subTests := []struct {
		name string
		//body        // whatever obj you need
		expectedCode int
	}{
		{
			name: "SuccessfulRequest",
			expectedCode: 200,
		},
	}
	for _, subTest := range subTests {
		t.Run(subTest.name, func(t *testing.T) {
			// bodyJson, _ := json.Marshal(subTest.body)
			// body := bytes.NewReader(credentials)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/tasks", nil) // body if exists
			tasksHandler(w, req)
			// res := w.Result() if you need data
			if subTest.expectedCode != w.Code {
				t.Errorf("expected status code %v, got status code %v", subTest.expectedCode, w.Code)
			}
		})
	}
}

var blood = TaskEntered{UserID: 1, TaskID: 1}
func TestPostTask(t *testing.T) {
	subTests := []struct {
		name string
		body  TaskEntered     // whatever obj you need
		expectedCode int
	}{
		{
			name: "SuccessfulCreateRequest",
			body: blood,
			expectedCode: 201,
		},
	}
	
	for _, subTest := range subTests {
		t.Run(subTest.name, func(t *testing.T) {
			bodyJson, _ := json.Marshal(subTest.body)
			body := bytes.NewReader(bodyJson)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/logtask", body) // body if exists
			postTask(w, req)
			// res := w.Result() // if you need data
			if subTest.expectedCode != w.Code {
				t.Errorf("expected status code %v, got status code %v", subTest.expectedCode, w.Code)
			}
		})
	}
}

func TestGetEnteredTasks(t *testing.T) {
	subTests := []struct {
		name string
		//body        // whatever obj you need
		expectedCode int
	}{
		{
			name: "SuccessfulRequest",
			expectedCode: 200,
		},
	}
	for _, subTest := range subTests {
		t.Run(subTest.name, func(t *testing.T) {
			// bodyJson, _ := json.Marshal(subTest.body)
			// body := bytes.NewReader(credentials)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/getenteredtasks", nil) // body if exists
			getEnteredTasks(w, req)
			// res := w.Result() if you need data
			if subTest.expectedCode != w.Code {
				t.Errorf("expected status code %v, got status code %v", subTest.expectedCode, w.Code)
			}
		})
	}
}

func TestGetFacilitiesDropdownDataByRegion(t *testing.T) {
	subTests := []struct {
		name string
		//body        // whatever obj you need
		expectedCode int
	}{
		{
			name: "SuccessfulRequest",
			expectedCode: 200,
		},
	}
	for _, subTest := range subTests {
		t.Run(subTest.name, func(t *testing.T) {
			// bodyJson, _ := json.Marshal(subTest.body)
			// body := bytes.NewReader(credentials)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/getfacilitiesdropdowndatabyregion", nil) // body if exists
			getFacilitiesDropdownDataByRegion(w, req)
			// res := w.Result() if you need data
			if subTest.expectedCode != w.Code {
				t.Errorf("expected status code %v, got status code %v", subTest.expectedCode, w.Code)
			}
		})
	}
}

var covid = Task{Name: "Covid Test", DepartmentID: 1}
func TestCreateNewTasks(t *testing.T) {
	subTests := []struct {
		name string
		body  Task     // whatever obj you need
		expectedCode int
	}{
		{
			name: "SuccessfulCreateRequest",
			body: covid,
			expectedCode: 201,
		},
	}
	
	for _, subTest := range subTests {
		t.Run(subTest.name, func(t *testing.T) {
			bodyJson, _ := json.Marshal(subTest.body)
			body := bytes.NewReader(bodyJson)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/createnewtask", body) // body if exists
			createNewTasks(w, req)
			// res := w.Result() // if you need data
			if subTest.expectedCode != w.Code {
				t.Errorf("expected status code %v, got status code %v", subTest.expectedCode, w.Code)
			}
		})
	}
}

func TestGetregionandfacilityHandler(t *testing.T) {
	subTests := []struct {
		name string
		//body        // whatever obj you need
		expectedCode int
	}{
		{
			name: "SuccessfulRequest",
			expectedCode: 200,
		},
	}
	for _, subTest := range subTests {
		t.Run(subTest.name, func(t *testing.T) {
			// bodyJson, _ := json.Marshal(subTest.body)
			// body := bytes.NewReader(credentials)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/getregionandfacility", nil) // body if exists
			getregionandfacilityHandler(w, req)
			// res := w.Result() if you need data
			if subTest.expectedCode != w.Code {
				t.Errorf("expected status code %v, got status code %v", subTest.expectedCode, w.Code)
			}
		})
	}
}
