package main

import "fmt"

func main() {
	fmt.Println("This is Test git-go")
	fmt.Println("This is Test git-go fix")
	sum := sum(5, 4)
	fmt.Printf("sum : %v", sum)
}

func sum(a, b int) int {
	return a + b
}
