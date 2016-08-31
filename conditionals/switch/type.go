package main

import "fmt"

func main() {
	switch "Tevin" {
	case "Daniel":
		fmt.Println("Hello Daniel")
	case "Tevin":
		fmt.Println("Hello Tevin")
	case "Jenny":
		fmt.Println("Hello Jenny")
	default:
		fmt.Println("You have no friends")
	}
}
