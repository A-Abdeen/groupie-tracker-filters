package gt

import (
	"html/template"
	"net/http"

	API "gt/webapp/API"
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
	members := r.FormValue("members")
	mincreation := r.FormValue("minCreationDate")
	maxcreation := r.FormValue("maxCreationDate")
	minAlbum := r.FormValue("minAlbumDate")
	maxAlbum := r.FormValue("maxAlbumDate")
	location := r.FormValue("locations")

	filteredDataToReturn, err := API.OrNotTosearch(members, mincreation, maxcreation, minAlbum, maxAlbum, location, MinAndMaxDatess, APIcall)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	var response API.Response
	response.Artists = filteredDataToReturn
	response.MinAndMaxDates = MinAndMaxDatess
	// fmt.Println(response.Artists)
	t.ExecuteTemplate(w, "base.html", response) // execution of all artists details to be presented in the homepage using base.html
}
