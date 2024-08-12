package main

import "fmt"

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
}

func getInfo(students []Student) {
	var bestMath, bestScience, bestEnglish, bestAverage Student
	highestMath, highestScience, highestEnglish, highestAverage := 0, 0, 0, 0

	for _, student := range students {
		if student.MathScore > highestMath {
			bestMath = student
			highestMath = student.MathScore
		}
		if student.ScienceScore > highestScience {
			bestScience = student
			highestScience = student.ScienceScore
		}
		if student.EnglishScore > highestEnglish {
			bestEnglish = student
			highestEnglish = student.EnglishScore
		}

		averageScore := (student.MathScore + student.ScienceScore + student.EnglishScore) / 3
		if averageScore > highestAverage {
			bestAverage = student
			highestAverage = averageScore
		}
	}

	fmt.Printf("best student in math: %s with %d\n", bestMath.Name, bestMath.MathScore)
	fmt.Printf("best student in science: %s with %d\n", bestScience.Name, bestScience.ScienceScore)
	fmt.Printf("best student in english: %s with %d\n", bestEnglish.Name, bestEnglish.EnglishScore)
	fmt.Printf("best student in average: %s with %d\n", bestAverage.Name, highestAverage)
}
