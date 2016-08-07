package main

import (
	"fmt"

	"github.com/tmantock/LearningGo/scope/package-scope/ex-visible/visibility"
)

func main() {
	fmt.Println(visibility.MyName)
	visibility.PrintVar()
}
