package gt

import (
	API "gt/webapp/API"
	"html/template"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	//fmt.Println("ErrorHandler is called.") // XXX
	var errorMsg string
	switch {
	case statusCode == 400:
		errorMsg = "Bad User Request"
	case statusCode == 404:
		errorMsg = "Page Not Found"
	case statusCode == 405:
		errorMsg = "Method Not Allowed"
	case statusCode == 413:
		errorMsg = "Limit Reached"
	case statusCode == 418:
		errorMsg = "The Server Refuses The Attempt To Brew Coffee With A Teapot"
	case statusCode == 500:
		errorMsg = "Internal Server Error"
	default:
		errorMsg = "UNFAMILIAR ERROR WHAT DIS :("
	}
	var errorResponse API.WebHandler
	errorResponse = errorResponse.PassError(errorMsg, statusCode)
	t, err := template.ParseFiles(HtmlTmpl...)
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "error.html", errorResponse)

}
