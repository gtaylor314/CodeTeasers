package main

import (
	"fmt"
	"math"
	"time"
)

// layout for golang's time package parse function
const timeForm = "15:04"

func main() {
	// Goal: calculate the smaller angle between the hour and minute hands of an analog clock
	// time is given in 24 format as hh:mm

	// test string to work with
	timeStr := "22:30"
	// timeStr := "00:30"
	// timeStr := "12:33"
	// timeStr := "01:59"
	// should produce errors
	// timeStr := "00:60"
	// timeStr := "15"
	// timeStr := "15:aa"
	// timeStr := "aa:30"
	// timeStr := "24:30"
	// parse the time string into a time.Time type variable
	time, err := time.Parse(timeForm, timeStr)
	if err != nil {
		fmt.Printf("Unable to parse time - error message: %s", err.Error())
		return
	}

	// given the time is in the 24 hour format, need to adjust the hour to the 12 hour format
	// the angle of the hour hand is the same whether the time is 01:00 or 13:00
	hour := float64(time.Hour())
	if hour > 12.0 {
		// converts the hour to the 12 hour format
		hour -= 12.0
	}
	fmt.Println(hour)

	minutes := float64(time.Minute())

	// calculate the angle of the hour hand - 12:00AM/PM represents 0 degrees
	// the hour hand travels a full 360 degress every 12 hours - i.e. 360/12 = 30 degrees/hour
	// in other words, the hour hand travels 360 degrees over 720 minutes - i.e 360/(12*60) = .5 degrees/min
	// using the hour and minutes we can calculate the angle of the hour offset by the minutes past the hour
	// i.e. h * (360/12) + m * (360/720)
	if hour == 12.0 {
		// adjust the hour to 0 if at 12:00PM
		// if at 12:00AM, the hour will already be 0 from the original 24 hour time format
		hour = 0.0
	}
	fmt.Println(hour)
	angleOfHour := (hour * (360.0 / 12.0)) + (minutes * (360.0 / 720.0))
	fmt.Println(angleOfHour)

	// calculate the angle of the minute hand
	// the minute hand travels a full 360 degrees every 60 minutes - i.e. 360/60 = 6 degrees/minute
	angleOfMinutes := minutes * (360.0 / 60.0)
	fmt.Println(angleOfMinutes)

	// calculate the angle between the hour hand and the minute hand
	// using absolute value we can always subtract the angle of the minute hand from the angle of the hour hand
	// if the result is greater than 180 degrees, we need to subtract it from 360 to get the smaller angle
	angle := math.Abs(angleOfHour - angleOfMinutes)
	if angle > 180.0 {
		angle = 360.0 - angle
	}

	// print resulting angle
	fmt.Printf("The shorter angle between the hour and minute hand for time %s is %f degrees", timeStr, angle)
}
