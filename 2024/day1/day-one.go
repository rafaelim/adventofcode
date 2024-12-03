package day1

import "fmt"

func newDayOne(handler DayOneHandler) *DayOne {
	return &DayOne{handler: handler}
}

func (d *DayOne) TestControlData() {
	distances, result := d.handler.TestControlData()
	fmt.Println("Distance list: ", distances)
	fmt.Println("Final Result: ", result)
}

func (d *DayOne) Run() {
	distances, result := d.handler.Run()
	fmt.Println("Distance list: ", distances)
	fmt.Println("Final Result: ", result)
}
