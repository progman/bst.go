package main

import (
 "fmt"
// "math"
)

func main() {

	var x int = 1;
	var y int = 2;


	var p_x *int = &x;
	var p_y *int = &y;


	fmt.Printf("x: %d\n", *p_x)
	fmt.Printf("y: %d\n", *p_y)


	p_x = p_y;


	fmt.Printf("x: %d\n", *p_x)
	fmt.Printf("y: %d\n", *p_y)
}
