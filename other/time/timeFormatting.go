// This program covers several basic functionality of time packages such as: 
// reading curent time in unix format, change time zone and printing different date-time format
package main

import (
	"fmt"
	"time"
)

func main() {
	_time := time.Now()

	fmt.Println("Time now:", _time)
	fmt.Println("Epoch time:", _time.Unix())

	fmt.Println("Day of the week is:", _time.Weekday())
	fmt.Println(_time.Format(time.RFC850))

	// Load time zone
	loc,_ :=  time.LoadLocation("Europe/Belgrade")

	locTime := _time.In(loc)
	fmt.Println("Belgrade:", locTime)
	
	// Custom format
	fmt.Println(locTime.Format("02.01.2006 15:04:05"))
}