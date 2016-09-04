package slice

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)
}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z)
}

//even though everything in go is pass-by-value. Slices are pass by reference since they are pointers a location in memory
