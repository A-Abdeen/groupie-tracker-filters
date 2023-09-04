package gt

import (
	// "fmt"
	API "gt/webapp/API"
	"html/template"
	"net/http"
	// "strings"
	// "strconv"
)

func FiltersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/filters/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	typedData := r.FormValue("filters")
	dataToReturn := API.Tosearch(typedData, APIcall)
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "base.html", dataToReturn)

}
