package main

import (
	"aoc/2024/day1"
	"flag"
	"log"
)

func main() {
	day := flag.Int("day", 1, "Challenge Day")
	part := flag.Int("part", 1, "Part of the day's challenge")
	test := flag.Bool("test", false, "Run with the test input")
	flag.Parse()

	switch *day {
	case 1:
		{
			day1.Run(*part, *test)
		}
	default:
		log.Fatal("\nNo handler find for the day %d", *day)
	}

}
