package main

import (
	"encoding/json"
	"fmt"
)

//Person struct defines a new person type
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James","Last":"Bond","Age":20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("-------------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
