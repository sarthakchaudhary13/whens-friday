package main

import (
	"fmt"
	"time"
)

func main() {
	whensFriday()
}

// WhensFriday will tell the user in how many days the comming
// Friday is or if it's the weekend already.
func whensFriday() {
	today := time.Now().Weekday()
	if today == time.Saturday || today == time.Sunday {
		fmt.Printf("It's %v, have fun!\nNext friday is in %v days.\n",today, nextFriday(today))
		return
	}

	switch time.Friday {
	case today + 0:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("Day after tomorrow.")
	default:
		fmt.Println("Too far away :(")
	}
}

// nextFriday will return in how many
// days the Friday will be.
func nextFriday(today time.Weekday) (nextFriday int) {
	nextFriday = int(time.Friday - today)
	if nextFriday < 0 {
		nextFriday += 7
	}
	return
}