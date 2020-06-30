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

	fmt.Printf("insert(%d)\n", key)
	var bstr_item bst_item_t
	bstr_item.init(key)


	if (p.head == nil) {

		fmt.Println("set as root")
		p.head = &bstr_item
		return
	}


	var p_cur *bst_item_t = p.head
	for {

// rigth way
		if (key > p_cur.key) {

			if (p_cur.r == nil) {
				p_cur.r = &bstr_item
				break;
			} else {
				p_cur = p_cur.r
			}
		}

// left way
		if (key < p_cur.key) {

			if (p_cur.l == nil) {
				p_cur.l = &bstr_item
				break;
			} else {
				p_cur = p_cur.l
			}
		}

// equal way (we can skip it or change), change it
		if (key == p_cur.key) {

			var p_old *bst_item_t = p_cur
			p_cur = &bstr_item
			bstr_item.l = p_old
			bstr_item.r = p_old.r
			p_old.r = nil
			break
		}
	}
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
