package main

//To receive one number
//Validate if the number is negative if it is to be transformed to negative (nun *-1)
//Save is one variable
//return

import (
	"fmt"
)

func main() {
	var number int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&number)
	fmt.Printf("You typed: %d\n", number)

	if number < 0 {
		fmt.Printf("The Number negativ is: %d", number)
	} else {
		number *= -1
		fmt.Printf("The negativ number is: %d", number)
	}
}
