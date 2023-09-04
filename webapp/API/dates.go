package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var allDates map[int][]string

func Dates() map[int][]string {
	fullJso, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullDatespage, err := io.ReadAll(fullJso.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualDates TmpDates //struct created to be able to unmarshal the full data for locations from above URL
	err2 := json.Unmarshal(fullDatespage, &individualDates)
	if err2 != nil {
		fmt.Print(err2)
	}
	var dates []string
	detailsPageDates := individualDates.Index
	allDates := make(map[int][]string, len(detailsPageDates))
	for i, indiv := range detailsPageDates {
		for _, data := range indiv.Dates {
			dates = append(dates, data)
		}
		allDates[i] = dates
		dates = []string{}
	}
	return allDates
}
