package main

import "fmt"

func main() {
	var st [50]string
	fmt.Println(st)
	for i := 65; i < 100; i++ {
		st[i-65] = string(i)
	}
	fmt.Println(st)
}
