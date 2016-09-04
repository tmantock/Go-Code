package value

import "fmt"

func main() {
	age := 42
	changeMe(age)
	fmt.Println(age)
}

func changeMe(z int) {
	fmt.Println(z)
	z = 47
	fmt.Println(z)
}
