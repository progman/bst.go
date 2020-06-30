package main

import (
 "fmt"
// "math"
)

type bst_item_t struct {

	key int
	l *bst_item_t
	r *bst_item_t
}

func (p *bst_item_t) init(key int) {

	p.key = key
	p.l   = nil
	p.r   = nil
}



type bst_t struct {
	head *bst_item_t
}

func (p *bst_t) init() {

	p.head = nil
}



func (p *bst_t) insert(key int) {

	fmt.Printf("insert(%d)\n", key)
	var p_bstr_item *bst_item_t = &bst_item_t{}
	p_bstr_item.init(key)


//	if (p.head == nil) {
//		fmt.Println("== nil")
//	} else {
//		fmt.Println("!= nil")
//	}


	if (p.head == nil) {

		fmt.Println("set as root")
		p.head = p_bstr_item

//		fmt.Printf("key %d\n", p_bstr_item.key)
//		fmt.Printf("key %d\n", p.head.key)


//	if (p.head == nil) {
//		fmt.Println("== nil")
//	} else {
//		fmt.Println("!= nil")
//	}

		return
	}




	var p_cur *bst_item_t = p.head
	for {

// rigth way
		if (key > p_cur.key) {

			if (p_cur.r == nil) {

				fmt.Println("set rigth")
				p_cur.r = p_bstr_item
				break;
			} else {

				fmt.Println("go to rigth")
				p_cur = p_cur.r
			}
		}

// left way
		if (key < p_cur.key) {

			if (p_cur.l == nil) {

				fmt.Println("set left")
				p_cur.l = p_bstr_item
				break;
			} else {

				fmt.Println("go to left")
				p_cur = p_cur.l
			}
		}

// equal way (we can skip it or change), change it
		if (key == p_cur.key) {

			fmt.Println("swap")
			var p_old *bst_item_t = p_cur
			p_cur = p_bstr_item
			p_bstr_item.l = p_old
			p_bstr_item.r = p_old.r
			p_old.r = nil
			break
		}
	}
}


func main() {

	var bst bst_t
	bst.init()


//	bst.insert(20)
//	bst.insert(20)

	bst.insert(100)

	bst.insert(110)
	bst.insert(90)

	bst.insert(120)
	bst.insert(80)

	bst.insert(130)
	bst.insert(70)


//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
	fmt.Println("ok")
}
