package gt

import (
	"errors"
	"fmt"
	handlers "gt/webapp/handlers"
	"log"
	"net/http"
)

func Server() {
	// Assigned portal
	const port = ":8080" // Change if needed
	// Serve the css directory to make it available for html
	styles := http.FileServer(http.Dir("../webapp/static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", styles))
	// Main Request Handlers
	http.HandleFunc("/", handlers.BaseHandler)
	http.HandleFunc("/search/", handlers.SearchHandler)
	http.HandleFunc("/filters/", handlers.FiltersHandler)
	http.HandleFunc("/details/", handlers.DetailsHandler)
	
	// http.HandleFunc("/search/", handlers.SearchHandler)
	// Run the server
	fmt.Println("http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("We closed the server? \n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		log.Fatal(err)
	}
}
