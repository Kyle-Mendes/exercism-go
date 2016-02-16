package main

import "fmt"

const TestVersion = 1

func main() {
	var (
		year int
		leap bool
	)

	fmt.Print("Input a year to check if it is a leap year! \n> ")
	fmt.Scanln(&year)

	leap = IsLeapYear(year)

	if leap {
		fmt.Printf("%v is a leap year!", year)
	} else {
		fmt.Printf("%v is not a leap year.", year)
	}

}

func IsLeapYear(year int) bool {

	if year%4 == 0 {
		if year%400 == 0 {
			return true
		} else if year%100 == 0 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
