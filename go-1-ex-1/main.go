package main

import "fmt"

func main() {
	var firstName string = "Marvin"
	var lastName string = "Hegi"

	var dayOfBirth int = 24
	var monthOfBirth int = 12
	var yearOfBirth int = 2007

	var numberOfSiblings int = 1

	var heightInMeters float32 = 1.80

	var zodiacSign string = "Steinbock"
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
