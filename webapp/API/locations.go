package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Locations() map[int][]string {
	fullJson, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullpage, err := io.ReadAll(fullJson.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualLocations TmpLocations //struct created to be able to unmarshal the full data for locations from above URL
	err2 := json.Unmarshal(fullpage, &individualLocations)
	if err2 != nil {
		fmt.Print(err2)
	}
	var locations []string // locations created so that to be able to add indivitualLocations to allLocations
	detailsPageLocations := individualLocations.Index
	allLocations := make(map[int][]string, len(detailsPageLocations))
	for i, indiv := range detailsPageLocations { // to go through the full locations data
		for _, data := range indiv.LocationsDetailed { // to edit and append the dates into an array
			data = strings.ReplaceAll(data, "-", ", ")
			data = strings.ReplaceAll(data, "_", " ")
			data = strings.Title(data)
			locations = append(locations, data)
		}
		allLocations[i] = locations // appending the data for all the artists in order
		locations = []string{}
	}
	return (allLocations)
}
