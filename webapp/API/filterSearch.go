package gt

import (
	// handlers "gt/webapp/handlers"
	//"debug/macho"

	"strconv"
	"strings"
)

func OrNotTosearch(members1 string, members2 string, members3 string, members4 string, members5 string, members6 string, members7 string, members8 string, mincreation string, maxcreation string, minAlbum string, maxAlbum string, location string, dates MinAndMaxDates, APIcall []Artists) ([]Artists, error) {
	// albumYear := 0
	// albumMonth := 0
	// albumDay := 0
	minAlbumInt := 0
	maxAlbumInt := 0
	mincreationYear := 0
	maxcreationYear := 0
	memberscheck := false
	members1Int := 0
	members2Int := 0
	members3Int := 0
	members4Int := 0
	members5Int := 0
	members6Int := 0
	members7Int := 0
	members8Int := 0
	var err error

	if members1 == "" && mincreation == dates.MinCreationDate && maxcreation == dates.MaxCreationDate && minAlbum == dates.MinAlbumDate && maxAlbum == dates.MaxAlbumDate && location == "none" && members2 == "" && members3 == "" && members4 == "" && members5 == "" && members6 == "" && members7 == "" && members8 == "" {
		return APIcall, nil
	}

	if location == "none" {
		location = ""
	}
	// fmt.Println(firstAlbum)

	if members1 != "" {
		members1Int, err = strconv.Atoi(members1)
		if err != nil {
			return nil, err
		}
		memberscheck = true
	}
	if members2 != "" {
		members2Int, err = strconv.Atoi(members2)
		if err != nil {
			return nil, err
		}
		memberscheck = true
	}
	if members3 != "" {
		members3Int, err = strconv.Atoi(members3)
		if err != nil {
			return nil, err
		}
		memberscheck = true
	}
	if members4 != "" {
		members4Int, err = strconv.Atoi(members4)
		if err != nil {
			return nil, err
		}
		memberscheck = true
	}
	if members5 != "" {
		members5Int, err = strconv.Atoi(members5)
		if err != nil {
			return nil, err
		}
		memberscheck = true
	}
	if members6 != "" {
		members6Int, err = strconv.Atoi(members6)
		if err != nil {
			return nil, err
		}
		memberscheck = true
	}
	if members7 != "" {
		members7Int, err = strconv.Atoi(members7)
		if err != nil {
			return nil, err
		}
		memberscheck = true
	}
	if members8 != "" {
		members8Int, err = strconv.Atoi(members8)
		if err != nil {
			return nil, err
		}
		memberscheck = true
	}

	if location != "" {
		location = strings.ToUpper(location)
	}
	// if firstAlbum != "" {
	// 	// Split the input string using "-"
	// 	albumDateComponents := strings.Split(firstAlbum, "-")

	// 	// Extract the components and convert them to integers
	// 	albumYear, err = strconv.Atoi(albumDateComponents[0])
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	albumMonth, err = strconv.Atoi(albumDateComponents[1])
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	albumDay, err = strconv.Atoi(albumDateComponents[2])
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// }

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

	if minAlbum != "" {
		// Extract the year and convert to int
		minAlbumInt, err = strconv.Atoi(minAlbum)
		if err != nil {
			return nil, err
		}
	}
	if maxAlbum != "" {
		// Extract the year and convert to int
		maxAlbumInt, err = strconv.Atoi(maxAlbum)
		if err != nil {
			return nil, err
		}
	}

	var dataToReturn []Artists
	ifMatching := false
	for i, oneArtist := range APIcall {
		if memberscheck {
			if len(oneArtist.Member) == members1Int || len(oneArtist.Member) == members2Int || len(oneArtist.Member) == members3Int || len(oneArtist.Member) == members4Int || len(oneArtist.Member) == members5Int || len(oneArtist.Member) == members6Int || len(oneArtist.Member) == members7Int || len(oneArtist.Member) == members8Int {
				ifMatching = true
			} else {
				ifMatching = false
				continue
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
		if minAlbum != "" && maxAlbum != "" {
			// Split the input string using "-"
			DateComponents := strings.Split(oneArtist.FirstAlbum, "-")

			// Extract the components and convert them to integers
			year := Atoi(DateComponents[2])
			// fmt.Println(DateComponents[2])

			// month := Atoi(DateComponents[1])
			// fmt.Println(DateComponents[1])

			// day := Atoi(DateComponents[0])
			// fmt.Println(DateComponents[0])

			// if albumYear < year {
			// 	ifMatching = true
			// } else if albumYear == year {
			// 	if albumMonth <= month {
			// 		if albumDay <= day {
			// 			ifMatching = true
			// 		} else {
			// 			ifMatching = false
			// 			continue
			// 		}
			// 	} else {
			// 		ifMatching = false
			// 		continue
			// 	}
			// } else {
			// 	ifMatching = false
			// 	continue
			// }

			if (minAlbumInt <= year) && (maxAlbumInt >= year) {
				ifMatching = true
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
			if ArtistDetails.Name == "Post Malone" {
				ArtistDetails.Name = "Yaman Almasri"
			}
			dataToReturn = append(dataToReturn, ArtistDetails)
		}
		ifMatching = false

	}
	return dataToReturn, err
}
