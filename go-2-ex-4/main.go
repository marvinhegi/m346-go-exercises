package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Students []Student
}

func main() {

	student1 := Student{FirstName: "Marvin", LastName: "Hegi"}
	student2 := Student{FirstName: "Nils", LastName: "Utiger"}
	student3 := Student{FirstName: "Yannick", LastName: "Blatty"}
	student4 := Student{FirstName: "Artur", LastName: "Ferreitra Cruz"}
	student5 := Student{FirstName: "Tobias", LastName: "Clausen"}

	class1 := Class{Students: []Student{student1, student2, student3}}
	class2 := Class{Students: []Student{student4, student5}}

	module := map[string][]Class{
		"346": {class1},
		"114": {class1, class2},
		"320": {class2},
	}

	for moduleID, classes := range module {
		fmt.Printf("Modul %s wird von diesen Klassen besucht:\n", moduleID)
		for i, class := range classes {
			fmt.Printf("  Klasse %d:\n", i+1)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.FirstName, student.LastName)
			}
		}
		fmt.Println()
	}
}
