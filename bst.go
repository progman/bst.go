/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
package main
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
import (
 "fmt"
// "math"
)
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
type bst_item_t struct {

	key int
	l *bst_item_t
	r *bst_item_t
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_item_t) init(key int) {

	p.key = key
	p.l   = nil
	p.r   = nil
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
type bst_t struct {
	head *bst_item_t
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_t) init() {

	p.head = nil
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
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

		fmt.Printf("\tset as root\n")
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


	if (key == p.head.key) {

		fmt.Printf("\thead swap\n")
		var p_old *bst_item_t = p.head
		p.head = p_bstr_item
		p_bstr_item.l = p_old
		p_bstr_item.r = p_old.r
		p_old.r = nil
		return
	}



	var p_cur *bst_item_t = p.head
	for {

// rigth way
		if (key > p_cur.key) {

			if (p_cur.r == nil) {

				fmt.Printf("\tset rigth\n")
				p_cur.r = p_bstr_item
				break;

			} else {

				fmt.Printf("\tgo to rigth\n")
				p_cur = p_cur.r
				continue;
			}
		}

// left way
		if (key < p_cur.key) {

			if (p_cur.l == nil) {

				fmt.Printf("\tset left\n")
				p_cur.l = p_bstr_item
				break;

			} else {

				fmt.Printf("\tgo to left\n")
				p_cur = p_cur.l
				continue;
			}
		}

// equal way (we can skip it or change), change it
		if (key == p_cur.key) {

			fmt.Printf("\tswap\n")
			var p_old *bst_item_t = p_cur
			p_cur = p_bstr_item
			p_bstr_item.l = p_old
			p_bstr_item.r = p_old.r
			p_old.r = nil
			break
		}
	}
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
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


	bst.insert(100)


	bst.insert(60)

	bst.insert(60)

	bst.insert(50)

	bst.insert(65)

	bst.insert(95)

//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
	fmt.Println("ok")
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
