package gt

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func LoadArtist() []Artists {

	fullJson, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fullArtistpage, err := io.ReadAll(fullJson.Body)
	if err != nil {
		log.Fatal(err)
	}
	var individualArtists []Artists //struct created to be able to unmarshal the full data for artists from above URL
	err2 := json.Unmarshal(fullArtistpage, &individualArtists)
	if err2 != nil {
		fmt.Print("Hello, World! ") // Indicator of partial marshalling success
	}
	return individualArtists // variable has all the artists data to be displayed in the homepage
}
