package routePack

import (
	//"context"
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

// func TestGetNurseFormHandler(t *testing.T) {
// 	response := tpl.ExecuteTemplate(w, "nurseLogin.html", nil)
// 	assert.NotNil(t, response, "The `response` should not be `nil`")
// }

func TestGetNurseFormHandler(t *testing.T) {
	r, _ := io.Pipe()
	_, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		t.Fatalf("failed to read template: %v", err)
	}
}

// func TestGetSupFormHandler(t *testing.T) {

// }

// func TestGetSupHomeHandler(t *testing.T) {

// }

// func TestGetNurseTasks(t *testing.T) {

// }
