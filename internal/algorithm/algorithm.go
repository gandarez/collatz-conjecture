package algorithm

// Calculate calculates 3N+1 problem and return an array of integers containing the calculated numbers.
func Calculate(input int) []int {
	var (
		results    []int
		calculated int
	)

	if input%2 == 0 {
		calculated = input / 2
	} else {
		calculated = input*3 + 1
	}

	results = append(results, calculated)

	if calculated == 1 {
		return results
	}

	results = append(results, Calculate(calculated)...)

	return results
}
