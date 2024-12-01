package dayone

import (
	"fmt"
	"log"
	"strconv"
)

func ExerciseTwo() {
	locationIDs := make(map[string]int)
	rows, err := getRows("./input/day_one_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	col1, col2, err := getSortedColumns(rows)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(col2); i++ {
		key := strconv.Itoa(col2[i])
		loc, ok := locationIDs[key]
		if !ok {
			locationIDs[key] = 1
			continue
		}
		locationIDs[key] = loc + 1
	}
	fmt.Println(locationIDs)

	sum := 0
	for i := 0; i < len(col1); i++ {
		key := strconv.Itoa(col1[i])

		if loc, ok := locationIDs[key]; ok {
			sum += col1[i] * loc
		}
	}

	fmt.Println(sum)
}
