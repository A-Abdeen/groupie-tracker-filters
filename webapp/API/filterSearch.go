package gt

import (
	// handlers "gt/webapp/handlers"
	//"debug/macho"

	"strconv"
	"strings"
)

func OrNotTosearch(members string, mincreation string, maxcreation string, firstAlbum string, location string, APIcall []Artists) ([]Artists, error) {
	albumYear := 0
	albumMonth := 0
	albumDay := 0
	mincreationYear := 0
	maxcreationYear := 0
	membersInt := 0
	var err error

	if members == "" && mincreation == "1900" && maxcreation == "2099" && firstAlbum == "" && location == "none" {
		return APIcall, nil
	}

	if location == "none" {
		location = ""
	}
	// fmt.Println(firstAlbum)

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

	if mincreation != "" {
		// Extract the year and convert to int
		mincreationYear, err = strconv.Atoi(mincreation)
		if err != nil {
			return nil, err
		}
	}
	if maxcreation != "" {
		// Extract the year and convert to int
		maxcreationYear, err = strconv.Atoi(maxcreation)
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
		if mincreation != "" && maxcreation != "" {
			if (mincreationYear <= oneArtist.Creationdate) && (maxcreationYear >= oneArtist.Creationdate) {
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
			// fmt.Println(DateComponents[2])

			month := Atoi(DateComponents[1])
			// fmt.Println(DateComponents[1])

			day := Atoi(DateComponents[0])
			// fmt.Println(DateComponents[0])

			if albumYear < year {
				ifMatching = true
			} else if albumYear == year {
				if albumMonth <= month {
					if albumDay <= day {
						ifMatching = true
					} else {
						ifMatching = false
						continue
					}
				} else {
					ifMatching = false
					continue
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
