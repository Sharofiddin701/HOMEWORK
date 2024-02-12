package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name   string
	Course int
}

func main() {
	students := []Student{
		{Name: "Shaxzod", Course: 1},
		{Name: "Lola", Course: 2},
		{Name: "Shaxzod", Course: 4},
		{Name: "Shaxzod", Course: 3},
		{Name: "Shaxzod", Course: 2},
		{Name: "Shaxzod", Course: 1},
	}


	sort.SliceStable(students, func(i, j int) bool {
		return students[i].Course < students[j].Course
	})


	for _, s := range students {
		fmt.Printf("Name: %s, Course: %d\n", s.Name, s.Course)
	}
}
