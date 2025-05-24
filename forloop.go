package main

import "fmt"

func main() {
	// basic syntax
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println(i)
	// }

	// classic for loop
	for j := 0; j <= 3; j++ {
		if j == 2 {
			continue
		} else if j == 3 {
			break
		}
		fmt.Println("j: ", j)
	}

	// for loop using range
	for k := range 6 {
		fmt.Println(k)
	}
}
