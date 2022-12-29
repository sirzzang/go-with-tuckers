package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	students := []Student{
		{"Sirzzang", 28},
		{"Chanbin", 19},
		{"Solbin", 25},
	}
	fmt.Println(students)
	sort.Sort(Students(students)) // ascending order
	fmt.Println(students)
	sort.Sort(sort.Reverse(Students(students))) // descending order
	fmt.Println(students)
}
