package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}

	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)

	sort.StringSlice(s).Sort()
	//sort.Sort(StringSlice(s))
	fmt.Println(s)

	//sort.String() takes a slice of strings
	sort.Strings(s)
	fmt.Println(s)

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	//Reverse string array
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	n := []int{5, 7, 8, 34, 67, 80, 10, 4, 11}

	fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

	sort.Ints(n)
	fmt.Println(n)
}
