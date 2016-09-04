package main

import "fmt"

func main() {
	fmt.Println(greet("Jane", "Doe"))
}

func greet(fname string, lname string) string {
	return fmt.Sprint(fname, lname)
}
