package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 42        //change the value at this address to 42
	fmt.Println(a) //42

	//b will be a pointer to the memory address where int a is stored
	//b is of type "int pointer"
	//*int -- the * is the part of the type -- b is of type *int
	//to see the value in that memory address, add a * in front of b (know as dereferencing)
	//everything in GO is pass by value
}
