package main

import "fmt"

func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	var y [256]int
	fmt.Println(len(y))
	fmt.Println(y[42])
	for i := 0; i < 256; i++ {
		y[i] = i
	}
	for i, v := range y {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}

	var w [256]byte
	fmt.Println(len(w))
	fmt.Println(w[42])
	for i := 0; i < 256; i++ {
		w[i] = byte(i)
	}
	for i, v := range w {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
