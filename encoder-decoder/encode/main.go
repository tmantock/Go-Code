package main

import (
	"encoding/json"
	"os"
)

//Person struct creates new Person type
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 22, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}
