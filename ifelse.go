package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Person is adult")
	} else {
		fmt.Println("Person is not adult")
	}

	role := "software developer"
	hasPermission := false

	if role == "software developer" && hasPermission {
		fmt.Println("Given Person is a", role, "and has a permission")
	} else {
		fmt.Println("Person has no persmission to enter inside the panel")
	}

}
