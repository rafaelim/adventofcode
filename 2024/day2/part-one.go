package day2

import (
	"fmt"
	"strconv"
)

func checkLevels(level []string, isIncreasing bool) bool {
	var a, b int
	switch {
	case len(level) == 0:
		return true
	case len(level) == 1:
		return true
	case len(level) > 1:
		{
			a, _ = strconv.Atoi(level[0])
			b, _ = strconv.Atoi(level[1])
		}
	}
	result := intAbs(a - b)
	if (result < 1 || result > 3) || (a < b && !isIncreasing) || (a > b && isIncreasing) {
		return false
	}

	return checkLevels(level[1:len(level)], isIncreasing)
}

func TestControlData() {
	reports := createList("day2/inputs/pt1-test")
	fmt.Println(reports)
	var safeLevels [][]string
	for _, levels := range reports {
		a, _ := strconv.Atoi(levels[0])
		b, _ := strconv.Atoi(levels[1])
		isSafe := checkLevels(levels, a < b)
		if isSafe {
			safeLevels = append(safeLevels, levels)
			fmt.Println("Safe: ", levels)
		} else {
			fmt.Println("Unsafe: ", levels)
		}
	}
	fmt.Println("Levels Safe: ", safeLevels)
	fmt.Println("Result: ", len(safeLevels))
}

func Run() {
	reports := createList("day2/inputs/input")
	fmt.Println(reports)
	var safeLevels [][]string
	for _, levels := range reports {
		a, _ := strconv.Atoi(levels[0])
		b, _ := strconv.Atoi(levels[1])
		isSafe := checkLevels(levels, a < b)
		if isSafe {
			safeLevels = append(safeLevels, levels)
			fmt.Println("Safe: ", levels)
		} else {
			fmt.Println("Unsafe: ", levels)
		}
	}
	fmt.Println("Levels Safe: ", safeLevels)
	fmt.Println("Result: ", len(safeLevels))
}
