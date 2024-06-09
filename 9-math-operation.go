package main 

import "fmt"

func main(){
	var a = 10
	var b = 5
	var c = a + b
	fmt.Println(c)

	fmt.Println(a * c)
	fmt.Println(c / b)
	fmt.Println(c - b)
	fmt.Println(c % a)

	// augmented operator
	a += 10
	fmt.Println(a)
	/* ada juga -= , *=, /= , %= */

	// unary operator
	b++
	fmt.Println(b)
	c--
	fmt.Println(c)
	/* ada juga - untuk menandakan negative number, + untuk positif contoh +100 (default number), dan ! untuk konver boolean value */
}