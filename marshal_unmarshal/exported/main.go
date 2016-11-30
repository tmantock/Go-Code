package main

import (
	"encoding/json"
	"fmt"
)

//Person struct defines a new person type
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	//Notice: notExported is not shown in string
	fmt.Println(string(bs))
}
