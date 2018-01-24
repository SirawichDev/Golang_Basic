package main

import "fmt"

func main() {
	var a int
	var b int
	var age int
	var name string
	fmt.Println("What's Your name ?")
	fmt.Scan(&name)
	fmt.Println("Hello! ", name)
	fmt.Println("How Old Are You?")
	fmt.Scan(&age)
	fmt.Println("Wow", name, " You're -> ", age, "Years old")
	fmt.Println("input number to a")
	fmt.Scan(&a)
	fmt.Println("input number to b")
	fmt.Scan(&b)
	fmt.Println("a+b is :", a+b)
	fmt.Println("a/b is :", a/b)
}
