/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
package main
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
import (
 "log"
// "fmt"
// "math"
)
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
//var list []*bst_item_t
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
type bst_item_t struct {

	parent *bst_item_t
	left   *bst_item_t
	right  *bst_item_t
	key int
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_item_t) init(key int) {

	p.parent = nil
	p.left   = nil
	p.right  = nil
	p.key    = key
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
func (p *bst_t) insert(key int, flag_uniq bool) {

	log.Printf("insert(%d)\n", key)
	var p_bstr_item *bst_item_t = &bst_item_t{}
//	list = append(list, p_bstr_item)


	p_bstr_item.init(key)


	if (p.head == nil) {

		log.Printf("\tset as root\n")
		p.head = p_bstr_item
		return
	}


// we will not use recursion (we think about stack size)
	var p_old *bst_item_t = nil
	var p_cur *bst_item_t = p.head
	for {

// rigth way
		if (key > p_cur.key) {

			if (p_cur.right == nil) {

				log.Printf("\tset rigth\n")
				p_bstr_item.parent = p_old
				p_cur.right        = p_bstr_item
				break;

			} else {

				log.Printf("\tgo to rigth\n")
				p_old = p_cur
				p_cur = p_cur.right
				continue;
			}
		}

// left way
		if (key < p_cur.key) {

			if (p_cur.left == nil) {

				log.Printf("\tset left\n")
				p_bstr_item.parent = p_old
				p_cur.left         = p_bstr_item
				break;

			} else {

				log.Printf("\tgo to left\n")
				p_old = p_cur
				p_cur = p_cur.left
				continue;
			}
		}

// equal way. we can skip it (use uniq values) or swap it (use non uniq values), swap by default
		if (key == p_cur.key) {

			if (flag_uniq == false) {

				log.Printf("\tswap\n")
				p_bstr_item.parent = p_old
				p_bstr_item.left   = p_cur.left
				p_cur.left         = p_bstr_item
				break

			} else {

				log.Printf("\tskip\n")
				break
			}
		}
	}
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_t) find(key int) *bst_item_t {

	log.Printf("find(%d)\n", key)


	if (p.head == nil) {

		log.Printf("\tbst is empty\n")
		return nil
	}


// we will not use recursion (we think about stack size)
	var p_cur *bst_item_t = p.head
	for {

// rigth way
		if (key > p_cur.key) {

			if (p_cur.right == nil) {

				log.Printf("\tdeath way\n")
				return nil;

			} else {

				log.Printf("\tgo to rigth\n")
				p_cur = p_cur.right
				continue;
			}
		}

// left way
		if (key < p_cur.key) {

			if (p_cur.left == nil) {

				log.Printf("\tdeath way\n")
				return nil;

			} else {

				log.Printf("\tgo to left\n")
				p_cur = p_cur.left
				continue;
			}
		}

// equal way
		if (key == p_cur.key) {

			log.Printf("\tequal\n")
			break
		}
	}


	return p_cur
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func main() {

	var bst bst_t
	bst.init()


	var flag_uniq bool = false


	bst.insert(100, flag_uniq)
	bst.insert(110, flag_uniq)
	bst.insert(90,  flag_uniq)
	bst.insert(100, flag_uniq)
	bst.insert(110, flag_uniq)
	bst.insert(120, flag_uniq)
	bst.insert(105, flag_uniq)


	var p *bst_item_t
	p = bst.find(90)
	if (p == nil) {

		log.Printf("is not found\n")

	} else {

		log.Printf("is found\n")
	}


	p = bst.find(100)
	if (p == nil) {

		log.Printf("is not found\n")

	} else {

		log.Printf("is found\n")
	}


	p = bst.find(120)
	if (p == nil) {

		log.Printf("is not found\n")

	} else {

		log.Printf("is found\n")
	}


	p = bst.find(777)
	if (p == nil) {

		log.Printf("is not found\n")

	} else {

		log.Printf("is found\n")
	}


	log.Printf("ok\n")


/*
	log.Printf("head: %p\n", bst.head)
	var i int
	for i=0; i < len(list); i++ {

		log.Printf("item:        %p\n", list[i])
		log.Printf("item.parent: %p\n", list[i].parent)
		log.Printf("item.left:   %p\n", list[i].left)
		log.Printf("item.right:  %p\n", list[i].right)
		log.Printf("item.key:    %d\n", list[i].key)

		log.Printf("\n")
	}
*/
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
