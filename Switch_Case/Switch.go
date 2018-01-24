package main

import "fmt"

func main() {

	switch "Alebrabra" {
	case "Sirawich":
		fmt.Println("Hey! Sirawich")
	case "Amagedon":
		fmt.Println("Hey! amagedon")
	case "Alebrabra":
		fmt.Println("Hey! Alebrabra")
		fallthrough
	case "Mommy":
		fmt.Println("Hey! Mommy")
	default:
		fmt.Println("No one")
	}
}
