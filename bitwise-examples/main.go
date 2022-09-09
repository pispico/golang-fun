package main

import (
	"fmt"
)

//https://medium.com/learning-the-go-programming-language/bit-hacking-with-go-e0acee258827
func main() {

	//example of bitwise operator &
	//Check if number is even or odd, it is faster than comparing with 'number % 2 == 0'
	for x := 0; x < 10; x++ {
		if x&1 == 1 {
			fmt.Println(x, " is Odd")
		} else {
			fmt.Println(x, " is Even")
		}
	}

}
