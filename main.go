package main

func findMaxInSeriesNumberofSlice(input []int) int {
	// var input = []int{7, 1, 2, 9, 7, 2, 1}
	var row []int
	var max = 0
	for a := range input {
		if a < len(input)-1 {
			if input[a+1]-1 == input[a] {
				if len(row) != 0 {
					if row[len(row)-1]+1 == input[a] {
						row = append(row, input[a])
					}
				} else {
					row = append(row, input[a])
				}
			} else {
				if len(row) != 0 {
					if input[a] == row[len(row)-1]+1 {
						row = append(row, input[a])
					} else {
						break
					}
				}
			}
		}
	}
	max = row[0]
	for b := range row {
		if max < row[b] {
			max = row[b]
		}
	}
	return max
}
