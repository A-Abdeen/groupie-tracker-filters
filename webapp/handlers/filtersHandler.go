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
	creation := r.FormValue("Creationdate")
	firstAlbum := r.FormValue("firstAlbumDate")
	locations := r.FormValue("concertLocations")

	filteredDataToReturn, err := API.OrNotTosearch(members, creation, firstAlbum,locations, APIcall)
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
	t.ExecuteTemplate(w, "base.html", filteredDataToReturn)
}
