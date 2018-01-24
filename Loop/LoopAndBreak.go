package main

import "fmt"

func main() {
	i := 0
	j := 0
	for {
		fmt.Println("x = ", i)
		fmt.Println("y = ", j)
		if i >= 10 {
			break
		}
		if j >= 8 {
			break
		}
		i++
		j++
	}
}
