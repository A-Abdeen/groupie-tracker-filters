package gt

import (
	API "gt/webapp/API"
	"html/template"
	"net/http"
	"strconv"
	// "fmt"
)

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("DetailsHandler is called.") // XXX
	// Verify Request Method
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	// Verify Request Pattern (Path)
	if r.URL.Path != "/details/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	idParam := r.URL.Query().Get("id")   // to get the ID of the band that was clicked at base.html
	idNumber, _ := strconv.Atoi(idParam) // convert the ID to an integer
	if idNumber <= 0 || idNumber > len(APIcall) {
		ErrorHandler(w, r, http.StatusNotFound) // since the ID can be entered using the url there should be a
		return                                  // condition for it being outside of the scope of the unmarshalled API
	}
	idNumber--                                                       // since the index starts from 0 and the id starts from 1, it needs to be adjusted to match
	DisplayDetails := APIcall[idNumber]                              // the unmarshalled APIcall of type []Artists can be indexed which defines DisplayDetails as type Artists
	DisplayDetails.Member = API.Arrangestring(DisplayDetails.Member) // arrange string created to be able to present the members better
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "details.html", DisplayDetails) // execution of each indivitual artist details to be presented in the display page using details.html
}
