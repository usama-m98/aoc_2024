package dayone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getRows(filePath string) ([]string, error) {
	rows := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("can't read:\n %v", err)
		return rows, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	return rows, scanner.Err()
}

func getSortedColumns(rows []string) ([]int, []int, error) {
	col1 := []int{}
	col2 := []int{}
	for _, col := range rows {
		str := strings.Split(col, "   ")
		intCol1, err := strconv.Atoi(str[0])
		if err != nil {
			log.Fatalf("can't convert to int:\n %v", err)
			return []int{}, []int{}, err
		}
		intCol2, err := strconv.Atoi(str[1])
		if err != nil {
			log.Fatalf("can't convert to int:\n %v", err)
			return []int{}, []int{}, err
		}
		col1 = append(col1, intCol1)
		col2 = append(col2, intCol2)
	}
	slices.Sort(col1)
	slices.Sort(col2)

	return col1, col2, nil
}

func ExerciseOne() {
	rows, err := getRows("./input/day_one_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	col1, col2, err := getSortedColumns(rows)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	for i := 0; i < len(col1); i++ {
		diff := col1[i] - col2[i]
		if diff < 0 {
			diff = -(diff)
		}
		sum += diff
	}

	fmt.Println(sum)
}
