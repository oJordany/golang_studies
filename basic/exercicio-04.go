package basic

func CalculateYears(years int) (result [3]int) {
	// for year := 1; year <= years; year++ {
	// 	result[0]++
	// 	if year == 1 {
	// 		result[1] = 15
	// 		result[2] = 15
	// 	} else if year == 2 {
	// 		result[1] += 9
	// 		result[2] += 9
	// 	} else {
	// 		result[1] += 4
	// 		result[2] += 5
	// 	}
	// }

	if years == 1 {
		return [3]int{1, 15, 15}
	}
	if years == 2 {
		return [3]int{2, 24, 24}
	}
	catYears := 24 + (years-2)*4
	dogYears := 24 + (years-2)*5
	result = [3]int{years, catYears, dogYears}
	return
}
