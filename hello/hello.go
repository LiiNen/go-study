package main

import "fmt"

func main() {
	const a int = iota
	const (
		b int = iota
		c int = iota
		d int = iota * iota
		e int = iota * iota * iota
	)

	fmt.Println(a, b, c, d, e)
}
