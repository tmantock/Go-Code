package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	//Overwrite
	//First string
	LicenseToKill bool
}

func (p person) greeting() {
	fmt.Println("Hello, I am a normal person.")
}

func (dz doubleZero) greeting() {
	fmt.Println("Hello I'm a secret agent.")
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		//First: "Joe",
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Sly",
			Last:  "Cooper",
			Age:   18,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.person.First, p1.LicenseToKill)
	fmt.Println(p2.First, p2.person.First, p2.LicenseToKill)

	p1.greeting()
	p1.person.greeting()
	p2.person.greeting()
	p2.greeting()
}
