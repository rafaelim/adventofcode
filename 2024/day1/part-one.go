package day1

func newPartOneHandler() *PartOneHandler {
	return &PartOneHandler{testInputPath: "day1/inputs/pt1-test", inputPath: "day1/inputs/input"}
}

func (h *PartOneHandler) getDistances(left, right []int) []int {
	var distances []int
	listLength := len(left)

	for range listLength {
		var leftMinValue, rightMinValue int
		leftMinValue, left = pullMinValue(left)
		rightMinValue, right = pullMinValue(right)

		distance := intAbs(rightMinValue - leftMinValue)
		distances = append(distances, distance)
	}

	return distances
}

func (h *PartOneHandler) TestControlData() ([]int, int) {
	left, right := createLists(h.testInputPath)

	distances := h.getDistances(left, right)
	total := sumValues(distances)

	return distances, total
}

func (h *PartOneHandler) Run() ([]int, int) {
	left, right := createLists(h.inputPath)

	distances := h.getDistances(left, right)
	total := sumValues(distances)

	return distances, total
}
