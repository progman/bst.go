/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
package main
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
import (
 "fmt"
// "math"
)
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
var list []*bst_item_t

/*
<-->tmp := []byte{}


<-->f, err := os.Open(filename)
<-->if err != nil {

<--><-->return "", err
<-->}
<-->defer f.Close()


<-->buf := make([]byte, 4096)


<-->for {

<--><-->n, err := f.Read(buf)
<--><-->if err == io.EOF {

<--><--><-->break
<--><-->}
<--><-->if err != nil {

<--><--><-->return "", err
<--><-->}


<--><-->for i := 0; i < n; i++ {

<--><--><-->tmp = append(tmp, buf[i])
<--><-->}
<-->}
*/



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
func (p *bst_t) insert(key int, flag_uniq bool) {

	fmt.Printf("insert(%d)\n", key)
	var p_bstr_item *bst_item_t = &bst_item_t{}
	list = append(list, p_bstr_item)


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


/*
	if (key == p.head.key) {

		fmt.Printf("\thead swap\n")
		var p_old *bst_item_t = p.head
		p.head = p_bstr_item
		p_bstr_item.l = p_old
		p_bstr_item.r = p_old.r
		p_old.r = nil
		return
	}
*/


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

// equal way. we can skip it (use uniq values) or swap it (use non uniq values), swap by default
		if (key == p_cur.key) {

			if (flag_uniq == false) {

				fmt.Printf("\tswap\n")


//				var p_old *bst_item_t = p_cur
//				p_cur = p_bstr_item
//				p_bstr_item.l = p_old
//				p_bstr_item.r = p_old.r
//				p_old.r = nil


				p_bstr_item.l = p_cur.l
				p_cur.l = p_bstr_item
				break

			} else {

				fmt.Printf("\tskip\n")
				break
			}
		}
	}
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func main() {

	var bst bst_t
	bst.init()

/*
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
*/

	var flag_uniq bool = true


	bst.insert(100, flag_uniq)
	bst.insert(110, flag_uniq)
	bst.insert(90,  flag_uniq)
	bst.insert(100, flag_uniq)
	bst.insert(110, flag_uniq)
	bst.insert(120, flag_uniq)
	bst.insert(105, flag_uniq)


//	v := Vertex{3, 4}
//	fmt.Println(v.Abs())
	fmt.Println("ok")



	fmt.Printf("head: %p\n", bst.head)
	var i int
	for i=0; i < len(list); i++ {

		fmt.Printf("item:     %p\n", list[i])
		fmt.Printf("item.key: %d\n", list[i].key)
		fmt.Printf("item.l:   %p\n", list[i].l)
		fmt.Printf("item.r:   %p\n", list[i].r)

		fmt.Printf("\n")
	}


}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
