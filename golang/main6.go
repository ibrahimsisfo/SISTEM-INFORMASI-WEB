package main

import "fmt"

func main() {
	a := 10 
	b := 3
	fmt.Println("Penjumlahan (a + b):", a+b) 
	fmt.Println("Modulus / Sisa Bagi (a % b):", a%b) 

	x := 5
	y := 7
	fmt.Println("x < y:", x < y)   
	fmt.Println("x == y:", x == y) 

	p := true
	q := false
	fmt.Println("p && q (AND):", p && q) 
	fmt.Println("p || q (OR):", p || q) 
	fmt.Println("!p (NOT):", !p)         
}