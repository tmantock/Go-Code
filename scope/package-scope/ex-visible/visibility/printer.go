package visibility

import "fmt"

//PrintVar function exported for main
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
}
