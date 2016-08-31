package main

import "fmt"

type Contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("Type not identified")
	}

}

func main() {
	SwitchOnType(7)
	SwitchOnType("Tevin")
	var t = Contact{"Good day mate", "Tevin"}
	SwitchOnType(t)
}
