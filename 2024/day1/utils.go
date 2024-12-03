package day1

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func createLists(path string) ([]int, []int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := strings.Join(strings.Fields(scanner.Text()), " ")
		arr := (strings.Split(val, " "))
		parsedLeftValue, errRight := strconv.Atoi(arr[0])
		if errRight != nil {
			log.Fatalf("Error converting left value to int {}", errRight)
		}
		parsedRightValue, errRight := strconv.Atoi(arr[1])
		if errRight != nil {
			log.Fatalf("Error converting right value to int {}", errRight)
		}
		left = append(left, parsedLeftValue)
		right = append(right, parsedRightValue)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return left, right
}

func intAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func pullMinValue(arr []int) (int, []int) {
	minValue := slices.Min(arr)
	index := slices.Index(arr, minValue)

	return minValue, slices.Delete(arr, index, index+1)
}

func sumValues(arr []int) int {
	var total int
	for _, val := range arr {
		total += val
	}

	return total
}
