package main

import "fmt"

func main() {
	var age int
	var name string
	fmt.Println("What's Your name ?")
	fmt.Scan(&name)
	fmt.Println("Hello! ", name)
	fmt.Println("How Old Are You?")
	fmt.Scan(&age)
	fmt.Println("Wow", name, " You're -> ", age, "Years old")

}
