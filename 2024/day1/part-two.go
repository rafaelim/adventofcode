package day1

import (
	"slices"
)

func newPartTwoHandler() *PartTwoHandler {
	return &PartTwoHandler{testInputPath: "day1/inputs/pt2-test", inputPath: "day1/inputs/input"}
}

func (h *PartTwoHandler) findOccurrences(arr []int, val int) int {
	if !slices.Contains(arr, val) {
		return 0
	}
	var occ int
	for _, arrval := range arr {
		if arrval == val {
			occ += 1
		}
	}
	return occ
}

func (h *PartTwoHandler) getDistances(left, right []int) []int {
	var distances []int
	for _, val := range left {
		occurrences := h.findOccurrences(right, val)
		distances = append(distances, occurrences*val)
	}

	return distances
}

func (h *PartTwoHandler) TestControlData() ([]int, int) {
	left, right := createLists(h.testInputPath)

	distances := h.getDistances(left, right)
	total := sumValues(distances)

	return distances, total
}

func (h *PartTwoHandler) Run() ([]int, int) {
	left, right := createLists(h.inputPath)

	distances := h.getDistances(left, right)
	total := sumValues(distances)

	return distances, total
}
