package main

import "fmt"

func main() {
	mynameis := "Ex"
	switch {
	case len(mynameis) == 2:
		fmt.Println("Your name have len equal 2")
		fallthrough
	case mynameis == "EX":
		fmt.Println("That's Right Your name is Ex XD")
	case mynameis == "ex":
		fmt.Println("No Way!!")
	default:
		fmt.Println("Don't You Know my name ? Fukk")

	}
}
