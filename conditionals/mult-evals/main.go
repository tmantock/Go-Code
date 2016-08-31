package main

import "fmt"

func main() {
	switch "Hennifer" {
	case "Daniel", "Hennifer":
		fmt.Println("Hello Daniel or Heniffer")
	case "Tevin":
		fmt.Println("Hello Tevin")
	case "Jenny":
		fmt.Println("Hello Jenny")
	default:
		fmt.Println("You have no friends")
	}
}
