package main

import "fmt"

func main() {
	// number1 := 17
	// number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string
	if number1, number2 := 17, 24; number1 > number2 {
		fmt.Println("Number1 is bigger by")
		fmt.Println(number1 - number2)
	} else if number2 > number1 {
		fmt.Println("Number2 is bigger by")
		fmt.Println(number2 - number1)
	} else {
		fmt.Println(resultMessage)
		fmt.Println("Number1 and number2 are equal.")
	}
}
