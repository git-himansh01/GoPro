package main

import "fmt"

func main() {
	var a int = 20
	b := 30
	var c int
	c = a + b
	fmt.Println(c)
	var (
		d = 25
		e = 25
	)
	fmt.Println(d + e)
	fmt.Printf("%T\n", a)
	fmt.Println("Hello, World!")
}
