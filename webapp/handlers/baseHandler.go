package gt

import (
	"html/template"
	"net/http"

	API "gt/webapp/API"
	// "fmt"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("BaseHandler is called.") // XXX
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	// typedData := r.FormValue("search")
	// API.SuggestionBox(typedData, APIcall)
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	var response API.Response
	response.Artists = APIcall
	response.MinAndMaxDates = MinAndMaxDatess
	// fmt.Println(response.Artists)
	t.ExecuteTemplate(w, "base.html", response) // execution of all artists details to be presented in the homepage using base.html
}
