package gt

import (
	"fmt"
	"strings"
	"time"

	API "gt/webapp/API"
)

/*
The function below is meant to load list of files-
that are parsed for use of html/template library-
at a global scope.
*/

var (
	HtmlTmpl        []string // global variables to be used by other functions
	APIcall         []API.Artists
	MinAndMaxDatess API.MinAndMaxDates
	AllLocations    API.AllLocations
)

func Init() {
	fmt.Println("Initializing Global Variable") // XXX
	HtmlTmpl = []string{
		"../webapp/static/base.html",
		"../webapp/static/details.html",
		"../webapp/static/error.html",
		// Add new html / template names here
	}
	fmt.Println("Global Variable initialized") // XXX
	APIcall = API.LoadArtist()                 // used to unmarshal full data into APIcall
	allLocations := API.Locations()            // used to unmarshal locations
	allDates := API.Dates()                    // used to unmarshal dates
	allRelations := API.Relations()            // used to unmarshal relations
	for i := range APIcall {                   // for loop to add data unmarshalled above into APIcall
		APIcall[i].Locations = allLocations[i]
		APIcall[i].Dates = allDates[i]
		APIcall[i].Relations = allRelations[i]
	}

	AllLocations.Locations = make([]string, 0, len(allLocations))
	for _, indexLocation := range allLocations {
		AllLocations.Locations = append(AllLocations.Locations, indexLocation...)
	}
	uniqueLocations := make(map[string]bool)
	for _, location := range AllLocations.Locations {
		uniqueLocations[location] = true
	}

	AllLocations.Locations = make([]string, 0, len(uniqueLocations))
	for location := range uniqueLocations {
		AllLocations.Locations = append(AllLocations.Locations, location)
	}

	// findings for min and max dates
	minCreation := time.Now().Year()
	maxCreation := 0

	for i := range APIcall { // for loop to add data unmarshalled above into APIcall

		if APIcall[i].Creationdate == minCreation || APIcall[i].Creationdate == maxCreation {
			continue
		}

		if APIcall[i].Creationdate < minCreation {
			minCreation = APIcall[i].Creationdate
		}
		if APIcall[i].Creationdate > maxCreation {
			maxCreation = APIcall[i].Creationdate
		}
	}
	MinAndMaxDatess.MinCreationDate = fmt.Sprint(minCreation)
	MinAndMaxDatess.MaxCreationDate = fmt.Sprint(maxCreation)

	// findings for min and max album dates
	minAlbumYear := time.Now().Year()
	maxAlbumYear := 0

	for _, oneArtist := range APIcall {
		// Split the date string into components
		DateComponents := strings.Split(oneArtist.FirstAlbum, "-")
		// Extract the components and convert them to integers
		year := API.Atoi(DateComponents[2])
		if year == minAlbumYear || year == maxAlbumYear {
			continue
		}

		if year < minAlbumYear {
			minAlbumYear = year
		}
		if year > maxAlbumYear {
			maxAlbumYear = year
		}
	}

	MinAndMaxDatess.MinAlbumDate = fmt.Sprint(minAlbumYear)
	MinAndMaxDatess.MaxAlbumDate = fmt.Sprint(maxAlbumYear)
}
