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

func Relations() map[int][]string {
	fullJso, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullRelationspage, err := io.ReadAll(fullJso.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualRelations TmpAllConRel //struct created to be able to unmarshal the full data for relations from above URL
	var dateAndLocation []string         // dateAndLocation created so that to be able to add indivitualRelations to DisplayDetails
	var oneDateandLocation string
	var dates string
	err2 := json.Unmarshal(fullRelationspage, &individualRelations)
	if err2 != nil {
		fmt.Print(err2)
	}
	detailsPageRelations := individualRelations.Index
	allRelations := make(map[int][]string, len(detailsPageRelations))
	for i, Relationdetails := range detailsPageRelations { // to go through all the relations and arrange them in order 
		for idx, data := range Relationdetails.Relation { // since detailsPageRelations is a map idx is the location and data is the dates for each location
			for idx, i := range data { // i represents each single date and idx is its index
				if idx == 0 {
					dates = dates  + i // the array of dates is converted to a string with the dates seperated by a coma
				} else {
					dates = i + ", " + dates
				}
			} // the following three lines are used to better represent each location
			idx = strings.ReplaceAll(idx, "-", ", ")
			idx = strings.ReplaceAll(idx, "_", " ")
			idx = strings.Title(idx)
			oneDateandLocation = idx + ": " + dates
			dateAndLocation = append(dateAndLocation, oneDateandLocation) // the dates and locations are appended into a single array of string
			dates = ""
		}
		allRelations[i] = dateAndLocation // the dates and locations for each artist are appended together 
		dateAndLocation = []string{}
	}

	return allRelations
}
