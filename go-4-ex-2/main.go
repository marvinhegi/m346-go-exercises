package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(sideA float64, sideB float64) float64 {
    
	c := math.Sqrt(math.Pow(sideA, 2) + math.Pow(sideB, 2))
	return c
	
}

func main() {
    
	fmt.Println(computeHypotenuse(30, 40))  //50
	fmt.Println(computeHypotenuse(20, 30))  //36.05551275463989
	fmt.Println(computeHypotenuse(10, 10))  //14.142135623730951

}
