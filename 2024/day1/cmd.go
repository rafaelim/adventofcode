package day1

import "log"

func Run(part int, test bool) {
	var service *DayOne
	switch part {
	case 1:
		{
			handler := newPartOneHandler()
			service = newDayOne(handler)
		}
	case 2:
		{
			handler := newPartTwoHandler()
			service = newDayOne(handler)
		}
	default:
		log.Fatal("Invalid part id, allowed values is 1 or 2")
	}

	if test {
		service.TestControlData()
		return
	}
	service.Run()
}
