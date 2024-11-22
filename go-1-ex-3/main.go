package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1 
	var when = time.Now()

	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	//Aufruf über Git Bash
	//go run main.go > eyes.txt 2> dice.log
} 
