package day2

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func createList(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reports [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := strings.Join(strings.Fields(scanner.Text()), " ")
		arr := (strings.Split(val, " "))
		reports = append(reports, arr)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return reports
}

func intAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
