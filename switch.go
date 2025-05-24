package main

import (
	"fmt"
	"time"
)

func main() {

	//1::  simple switch command
	i := 1
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// 2:: Multiple condition switch

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's Weekend")
	default:
		fmt.Println("it's working day")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an Integer")
		case string:
			fmt.Println("It's a string")
		case bool:
			fmt.Println("It's a Boolean")
		default:
			fmt.Println("Other: ", t)
		}
	}

	whoAmI("golang")
	whoAmI(1)

}
