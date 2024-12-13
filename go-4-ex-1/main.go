package main

import (
	"errors"
	"fmt"
)

func berechneNote(erreichtePunkte, maxPunkte float64) (float64, error) {
    
	if erreichtePunkte < 0 || maxPunkte <= 0 {
		return 0, errors.New("Error: Punkte enthalten ungültige werte")
	}
	
	if erreichtePunkte > maxPunkte {
		return 0, errors.New("Error: erreichte Punkte sind größer als maximum")
	}
	
	note := 1 +( 5*(erreichtePunkte/maxPunkte))
	return note, nil
}

func main() {

	note1, error1 := berechneNote(8, 10.0)
	
	if error1 != nil {
		fmt.Println("Fehler:", error1)
	} else {
		fmt.Println("Note:", note1)  // Note 5
	}


	note2, error2 := berechneNote(10.0, 10.0)
	
	if error2 != nil {
		fmt.Println("Fehler:", error2)
	} else {
		fmt.Println("Note:", note2)   // Note 6
	}


	note3, error3 := berechneNote(-1.0, 10.0)
	
	if error3 != nil {
		fmt.Println("Fehler:", error3)  //error weil Punktzahl - ist
	} else {
		fmt.Println("Note:", note3)  
	}
}
