// Created By: Charlotte Jhu
// Created On: May 2023
//
// This function shows the user if they get free admission or not

package main

import (
	"fmt"
)

func main() {
	// This function shows the user if they get free admission or not
	// variables
	var userAge int
	var dayOfTheWeek string

	// input
	fmt.Println("Enter you age and the day of the week below to see if you get free admission.")
	fmt.Println("")
	fmt.Print("Enter your age: ")
	fmt.Scanln(&userAge)
	fmt.Print("Enter the day of the week: ")
	fmt.Scanln(&dayOfTheWeek)

	// process
	if (dayOfTheWeek == "Monday" || dayOfTheWeek == "Thursday") || (userAge > 12 && userAge < 21) {
		// output
		fmt.Println("You get free admission!")
	} else {
		// output
		fmt.Println("You have to pay admission.")
	}
	fmt.Println("\nDone.")
}
