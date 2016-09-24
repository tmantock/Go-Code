package main

import "fmt"

//Testing Go after macos Sierra upgrade

func main() {
	var x int
	var y int
	var sum int
	string := "Hello Mate! "
	stringToConcat := "Welcome to Australia"

	x = 0
	y = 1

	for i := 0; i < 100; i++ {
		sum += x + y
		fmt.Println(string + stringToConcat + " the new number is ")
		fmt.Println(sum)
		x++
		y++
	}
}
