package main

import (
	"encoding/json"
	"fmt"
)

//Person struct defines a new person type
type Person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	//Notice: how an empty JSON object is printed since nothing is exorted from the struct
	fmt.Println(string(bs))
}
