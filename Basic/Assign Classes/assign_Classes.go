package main

import (
	"fmt"
	"sort"
)

func main() {
	totalTestCases := 0
	fmt.Scan(&totalTestCases)
	var answers []int

	for i := 0; i < totalTestCases; i++ {
		answers = append(answers, 0)
	}

	for i := 0; i < totalTestCases; i++ {
		students, skillSet := Input()
		AssignClasses(students, skillSet, answers, i)

	}

	for i := 0; i < totalTestCases; i++ {
		fmt.Println(answers[i])
	}
}

func Input() (int, []int) {
	totalStudents := 0
	fmt.Scan(&totalStudents)

	var totalSkills []int
	for i := 0; i < 2*totalStudents; i++ {
		a := 0
		fmt.Scan(&a)
		totalSkills = append(totalSkills, a)
	}
	return totalStudents, totalSkills
}

func AssignClasses(students int, skillSet []int, answers []int, testcaseIdx int) {
	sort.Ints(skillSet)

	answer := skillSet[students] - skillSet[students-1]
	if answer < 0 {
		answer = -answer
	}
	answers[testcaseIdx] = answer
}
