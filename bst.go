package main

import (
 "fmt"
// "math"
)
/*
type Vertex struct {
 X, Y float64
}

func (v Vertex) Abs() float64 {
 return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
*/

type bst_item_t struct {

	key int
	l *bst_item_t
	r *bst_item_t
}

func (p bst_item_t) init(key int) {

	p.key = key
	p.l   = nil
	p.r   = nil
}



type bst_t struct {
	head *bst_item_t
}

func (p bst_t) init() {

	p.head = nil
}



func (p bst_t) insert(key int) {

	var bstr_item bst_item_t
	bstr_item.init(key)

	if (p.head == nil) {

		p.head = &bstr_item
		return
	}
/*
	for {

		if p.head.key < val
	}
*/
}


func main() {

	var bst bst_t
	bst.init()


	bst.insert(1)
	bst.insert(2)
	bst.insert(3)


//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
	fmt.Println("ok")
}
