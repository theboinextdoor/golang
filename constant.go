package main

import "fmt"

// const variable can be declared and initialize outside the function
const age = 25
const Totalpackage = 2500

func main() {

	// const value can never be changed once it is initialize
	const name = "Damon Salvatore"
	fmt.Println("My full name is : ", name)
	fmt.Println("Age is : ", age)
	fmt.Println("Package is : ", Totalpackage)

	// grouping variable constant
	// NOTE:: you cnnot change the value of it
	const (
		port   = 500
		host   = "localhost"
		domain = "my-Golang"
	)

	fmt.Println("Domain", domain, "is currently hosted", host, port)

}
