package router

import (
	"NursesTaskAnalysis/dbconnection"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

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

// package routePack

// import (
// 	"testing"
// 	"net/http/httptest"
// 	"NurseTasks/dbconnection"

// )

// func TestMain(m *testing.M) {
// 	setupTest()
// 	code := m.Run()
// 	// teardown
// 	os.Exit(code)
//   }

// func setupTest() {
// 	driver := dbconnection.Postgres
// 	dbconnection.New(driver)
// }

// type Tests struct {
// 	name string
// 	server *httptest.Server
// 	response *Task
// 	expectedError error
// }

// func TestGetEnteredTasks(t *testing.T) {

// 	tests := []Tests {
// 		{
// 		name: "get-tasks",
// 		server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Response){
// 			w.WriteHeader(http.StatusOK)
// 			w.Write([]byte(`{ "Name": "Covid test", "DepartmentID": 1  }`))
// 		})),
// 		response: &Task{
// 			Name: "Covid test",
// 			DepartmentID: 1,
// 		},
// 		expectedError: nill,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T){
// 			defer test.server
// 		})
// 	}
// }
