package main

import "fmt"

func main() {
	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Friend's name is of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Hello Tim")
	case myFriendsName == "Jerry":
		fmt.Println("Hello Jerry")
	case myFriendsName == "Tom":
		fmt.Println("Hello Tom")
	default:
		fmt.Println("No matches")
	}
}
