package main

import (
 "fmt"
// "math"
)


type item1_t struct {

	key int
}
func (p item1_t) init(key int) {

	p.key = key
}


type item2_t struct {

	key int
}
func (p *item2_t) init(key int) {

	p.key = key
}


func main() {

	var item1 item1_t
	item1.init(2020)
	fmt.Printf("key: %d\n", item1.key)

	var item2 item2_t
	item2.init(2020)
	fmt.Printf("key: %d\n", item2.key)
}
