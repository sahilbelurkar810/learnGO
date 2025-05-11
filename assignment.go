package main

import "fmt"

func main() {
	// checkHeight()
	// concat("Hello", "World")
	// shorthandConcat("Hello", "World")
	// sendSoFar := 430
	// const sendsToAdd = 25
	// sendSoFar = incrementSends(sendSoFar, sendsToAdd)
	// fmt.Println("Total sends:", sendSoFar)
	// firstName, _ := getNames()
	// fmt.Println("First Name:", firstName)
	// test(4)
	// test(18)
	// test(21)
	// test(65)

	quotient, remainder := divide(10, 0)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

}

func divide(dividend, divisor int) (int, int) {
	if divisor == 0 {
		fmt.Println("Cannot divide by zero")
		return 0, 0
	}
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// func getNames() (string, string) {
// 	return "John", "Doe"
// }

// func yearsUntilEvents(age int) (int, int, int) {

// 	yearsUntilAdult := 18 - age
// 	if yearsUntilAdult < 0 {
// 		yearsUntilAdult = 0
// 	}
// 	yearsUntilDrinking := 21 - age
// 	if yearsUntilDrinking < 0 {
// 		yearsUntilDrinking = 0
// 	}
// 	yearsUntilRetirement := 65 - age
// 	if yearsUntilRetirement < 0 {
// 		yearsUntilRetirement = 0
// 	}
// 	return yearsUntilAdult, yearsUntilDrinking, yearsUntilRetirement
// }

// func test(age int) {
// 	fmt.Println("===================================")
// 	fmt.Println("Age:", age)
// 	yearsUntilAdult, yearsUntilDrinking, yearsUntilRetirement := yearsUntilEvents(age)
// 	fmt.Println("Years until adult:", yearsUntilAdult)
// 	fmt.Println("Years until drinking:", yearsUntilDrinking)
// 	fmt.Println("Years until retirement:", yearsUntilRetirement)
// 	fmt.Println("===================================")
// }

// func incrementSends(currentSends int, sendsToAdd int) int {
// 	currentSends += sendsToAdd
// 	return currentSends
// }

// func checkHeight() {
// 	height := 6

// 	if height > 6 {
// 		fmt.Println("You are super tall enough to ride!")
// 	} else if height > 4 {
// 		fmt.Println("You are tall enough to ride!")
// 	} else if height > 2 {
// 		fmt.Println("You are tall enough to ride, but just barely!")
// 	} else {
// 		fmt.Println("You are not tall enough to ride!")
// 	}
// }

// func concat(x string, y string) {
// 	fmt.Println(x + " " + y)
// }

// func shorthandConcat(x, y string) string {
// 	result := x + " " + y
// 	fmt.Println(result)
// 	return result
// }
