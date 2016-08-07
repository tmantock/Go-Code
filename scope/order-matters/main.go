package main

import "fmt"

func main() {
	//x is undefined at this point order matters
	fmt.Println(x)
	//y is defined because of package scope
	fmt.Println(y)
	x := 42
}

var y = 42
