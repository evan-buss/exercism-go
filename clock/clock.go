// package main

// TODO: Finish this problem

package clock

import (
	"fmt"
	"math"
	"strconv"
)

// Clock is an object containing an hour and minute
type Clock struct {
	hour, minute int
}

// func main() {
// 	clock1 := New(1, -160)
// 	fmt.Println(clock1.String())
// }

// New creates a new clock object with the given hour and minute
func New(hour, minute int) Clock {
	newClock := Clock{hour: 0, minute: 0}

	// Set the hour
	if hour <= 24 && hour >= 0 {
		newClock.hour = hour
	} else if hour > 24 {
		newClock = newClock.Add(hour * 60)
	} else if hour < 0 {
		newClock = newClock.Subtract(hour * 60 * -1)
	}

	// Set the minutes
	if minute <= 59 && minute >= 0 {
		newClock.minute = minute
	} else if minute > 59 {
		newClock = newClock.Add(minute)
	} else if minute < 0 {
		newClock = newClock.Subtract(minute * -1)
	}

	return newClock
}

// Add increments the clock by the given minutes
func (clock Clock) Add(inc int) Clock {
	hour := clock.hour
	minute := clock.minute

	if minute+inc < 60 {
		return Clock{hour, minute + inc}
	}

	// Add hours the existing
	hour += int(math.Floor(float64((minute + inc) / 60)))
	// New minutes set to remainder
	minute = (minute + inc) % 60
	// Continually subtract 24 hours until hour less than 24
	for hour > 24 {
		hour = hour - 24
	}

	return Clock{hour, minute}
}

// Subtract decrements the clock by the given minutes
func (clock Clock) Subtract(dec int) Clock {
	hour, minute := clock.hour, clock.minute

	// Decrement less than the hour. Simple subtraction
	if minute-dec >= 0 {
		return Clock{hour, minute - dec}
	}
	// Otherwise perform hour shifts

	// decAmount := int(math.Floor(float64((minute - dec) / 60)))
	decAmount := float64(minute-dec) / 60
	fmt.Println("Calculated Dec: ", decAmount)
	if math.Floor(math.Abs(float64(decAmount))) == 0 {
		fmt.Println("floor 0")
		hour--
	} else {
		hour -= int(math.Floor(math.Abs(float64(decAmount))))
	}
	// hour -= int(math.Floor(float64((minute - dec) / 60)))

	// New minutes set to remainder
	// minute = int(math.Abs(float64(minute-dec))) % 60
	if minute == 0 {
		fmt.Println("minute is 0, therfore", 60-dec)
		// minute = (60 - dec) % 60
		minute = int(math.Abs(float64(60-dec))) % 60
	} else {
		minute = int(math.Abs(float64(minute-dec))) % 60
	}
	// Continually subtract 24 hours until hour less than 24
	for hour < 0 {
		hour = hour + 24
	}

	// fmt.Println("returning: ", Clock{hour, minute}.String())
	return Clock{hour, minute}
}

// String returns the clock's time in string format
func (clock Clock) String() string {
	hourStr, minStr := strconv.Itoa(clock.hour), strconv.Itoa(clock.minute)
	if clock.hour < 10 {
		hourStr = "0" + hourStr
	} else if clock.hour == 24 {
		hourStr = "00"
	}

	if clock.minute < 10 {
		minStr = "0" + minStr
	}
	return hourStr + ":" + minStr
}
