package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type foo int

func main() {
	p1 := person{"James", "Bond", 25}
	p2 := person{"Sly", "Cooper", 18}

	var myage foo
	myage = 23
	fmt.Printf("%T %v \n", myage, myage)

	var yourage int
	yourage = 44
	fmt.Printf("%T %v \n", yourage, yourage)

	//This will not work
	//fmt.Println(myage + yourage)

	//This will work
	fmt.Println(int(myage) + yourage)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
