package gt

import "strconv"

func Arrangestring(array []string) []string {
	var Member []string
	for i := 0; i < len(array); i++ {

		Membermiddles := "(" + strconv.Itoa(i+1) + ") " + array[i] + "  "
		Member = append(Member, Membermiddles)

	}
	array = Member
	return array
}
