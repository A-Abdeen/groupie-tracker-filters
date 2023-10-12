package gt

import (
	// handlers "gt/webapp/handlers"
	//"debug/macho"
	"strconv"
	"strings"
)

func OrNotTosearch(members string, creation string, firstAlbum string, location string, APIcall []Artists) ([]Artists, error) {
	albumYear := 0
	albumMonth := 0
	albumDay := 0
	creationYear := 0
	membersInt := 0
	var err error

	if members == "" && creation == "" && firstAlbum == "" && location == "none" {
		return APIcall, nil
	}

	if location == "none" {
		location = ""
	}

	if members != "" {
		membersInt, err = strconv.Atoi(members)
		if err != nil {
			return nil, err
		}
	}
	if location != "" {
		location = strings.ToUpper(location)
	}
	if firstAlbum != "" {
		// Split the input string using "-"
		albumDateComponents := strings.Split(firstAlbum, "-")

		// Extract the components and convert them to integers
		albumYear, err = strconv.Atoi(albumDateComponents[0])
		if err != nil {
			return nil, err
		}
		albumMonth, err = strconv.Atoi(albumDateComponents[1])
		if err != nil {
			return nil, err
		}
		albumDay, err = strconv.Atoi(albumDateComponents[2])
		if err != nil {
			return nil, err
		}

	}

	if creation != "" {
		// Extract the year and convert to int
		creationYear, err = strconv.Atoi(creation)
		if err != nil {
			return nil, err
		}
	}

	var dataToReturn []Artists
	ifMatching := false
	for i, oneArtist := range APIcall {
		if members != "" {
			if membersInt > 0 && membersInt < 8 {
				if len(oneArtist.Member) == membersInt {
					ifMatching = true
				} else {
					ifMatching = false
					continue
				}
			}
		}
		if creation != "" {
			if creationYear <= oneArtist.Creationdate {
				ifMatching = true
			} else {
				ifMatching = false
				continue
			}
		}
		if firstAlbum != "" {
			// Split the input string using "-"
			DateComponents := strings.Split(oneArtist.FirstAlbum, "-")

			// Extract the components and convert them to integers
			year := Atoi(DateComponents[2])

			month := Atoi(DateComponents[1])

			day := Atoi(DateComponents[0])

			if albumYear <= year {
				if albumMonth <= month {
					if albumDay <= day {
						ifMatching = true
					}
				}
			} else {
				ifMatching = false
				continue
			}
		}
		if location != "" {
			for _, oneLocation := range oneArtist.Locations {
				location = strings.ReplaceAll(location, "-", ", ")
				location = strings.ReplaceAll(location, "_", " ")
				if strings.ToUpper(oneLocation) == location {
					ifMatching = true
					break
				} else {
					ifMatching = false
					continue
				}
			}
		}
		if ifMatching {
			ArtistDetails := APIcall[i]
			dataToReturn = append(dataToReturn, ArtistDetails)
		}
		ifMatching = false
	}
	return dataToReturn, err
}
