package main

import (
 "fmt"
// "math"
)


type item_t struct {

	key int
}

func (p *item_t) init(key int) {

	p.key = key
}


func main() {

	var item item_t
	item.init(2020)

	fmt.Printf("key: %d\n", item.key)
}
