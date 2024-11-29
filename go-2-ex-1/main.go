package main

import (
	"fmt"
	"time"
)

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month time.Month
	Year  int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName:         FullName{"Marvin", "Hegi"},
		BirthDate:        BirthDate{24, 12, 2007},
		NumberOfSiblings: 1,
		ZodiacSign:       '\u2651',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
