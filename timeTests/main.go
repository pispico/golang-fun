package main

import (
	"fmt"
	"time"
)

func main() {

	//Parsing date
	date1 := "7/25/2019 13:45:00"
	t, err := time.Parse("1/02/2006 15:04:05", date1)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	//Comparing dates, check if date2 is after or before than current time
	date2 := "October 3, 2022 20:32:00"
	t2, err2 := time.Parse("January 2, 2006 15:04:05", date2)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(t2.Before(time.Now()))
	fmt.Println(t2.After(time.Now()))

	// Returns the Hour of a date3, the same can be done with the other attributes
	date3 := "Thursday, July 25, 2019 13:45:00"
	t3, err3 := time.Parse("Monday, January 2, 2006 15:04:05", date3)
	if err3 != nil {
		panic(err3)
	}
	t3h := t3.Hour()
	fmt.Println(t3h)

	// Description returns a formatted string of the time.
	date4 := "6/6/2005 10:30:00"
	t4, err4 := time.Parse("1/2/2006 15:04:05", date4)
	if err != nil {
		panic(err4)
	}
	fmt.Printf("You have an appointment on %s\n", t4.Format("Monday, January 2, 2006, at 15:04."))
}
