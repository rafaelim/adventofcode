package day1

type DayOne struct {
	handler DayOneHandler
}

type Handler struct {
	testInputPath string
	inputPath     string
}

type PartOneHandler Handler
type PartTwoHandler Handler

type DayOneHandler interface {
	getDistances(left, right []int) []int
	TestControlData() ([]int, int)
	Run() ([]int, int)
}
