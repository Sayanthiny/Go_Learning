package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Printf("enter the number :")
	fmt.Scanf("%d", &num)
	// n := num / 2
	// if num%2 == 0 {
	// 	fmt.Printf("%d is a even, %d times", num, n)
	// } else {
	// 	fmt.Printf("%d is a odd %d times ", num, n)
	// }
	fmt.Printf("%s", fizzBuzz(num))

}

func fizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return ("fizzBuzz")
	} else if n%3 == 0 {
		return ("fizz")
	} else if n%5 == 0 {
		return ("buzz")
	} else {
		return strconv.Itoa(n)
	}

}
