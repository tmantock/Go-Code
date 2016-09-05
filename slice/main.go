package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11, 13}
	fmt.Printf("%T \n", mySlice)
	fmt.Println(mySlice)

	myMakeSlice := make([]int, 0, 5)
	fmt.Println("________________________")
	fmt.Println(myMakeSlice)
	fmt.Println(len(myMakeSlice))
	fmt.Println(cap(myMakeSlice))
	fmt.Println("_________________________")

	for i := 0; i < 80; i++ {
		myMakeSlice = append(myMakeSlice, i)
		fmt.Println("Len: ", len(myMakeSlice), "Capacity: ", cap(myMakeSlice), "Value: ", myMakeSlice[i])
	}
}
