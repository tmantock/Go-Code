package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

//Person creates a new person type
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"James","Last":"Bond","Age":22}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
