package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
	//i is undefined when trying to print outside of the loop
	//fmt.Println(i)
	//See if i is within the scope of the loo[ or function]
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
