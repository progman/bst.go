/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
package main
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
import (
	"log"
	"sync"
//	"errors"

	"fmt"
	"net/http"
//	"html"
	"strconv"
)
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
//var list []*bst_item_t
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
type bst_item_t struct {

	parent *bst_item_t
	left   *bst_item_t
	right  *bst_item_t
	key int64
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_item_t) init(key int64) {

	p.parent = nil
	p.left   = nil
	p.right  = nil
	p.key    = key
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
type bst_t struct {

	head  *bst_item_t
	mutex sync.RWMutex
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_t) init() {

	p.head = nil
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_t) insert(key int64, flag_uniq bool) (*bst_item_t) {

	p.mutex.Lock()
	log.Printf("insert(%d)\n", key)
	var p_bstr_item *bst_item_t = &bst_item_t{}
//	list = append(list, p_bstr_item)


	p_bstr_item.init(key)


	if (p.head == nil) {

		log.Printf("\tset as root\n")
		p.head = p_bstr_item
		p.mutex.Unlock()
		return p_bstr_item
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


	p.mutex.Unlock()
	return p_bstr_item
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_t) findInner(key int64) (*bst_item_t) {

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
func (p *bst_t) find(key int64) (*bst_item_t) {

	p.mutex.RLock()
	log.Printf("find(%d)\n", key)


	p_cur := p.findInner(key)


	p.mutex.RUnlock()
	return p_cur
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_t) deleteInner(key int64) (*bst_t) {

	var p_cur *bst_item_t = p.findInner(key)
	if (p_cur == nil) {

		return nil
	}




//	parent *bst_item_t
//	left   *bst_item_t
//	right  *bst_item_t
//	key int64


	if (p_cur.left == nil) && (p_cur.right == nil) {

		if (p_cur.parent.left == p_cur) {
			p_cur.parent.left = nil
		}

		if (p_cur.parent.right == p_cur) {
			p_cur.parent.right = nil
		}

		return p
	}


//return errors.New("invalid hex length")


	return p
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func (p *bst_t) delete(key int64, flag_uniq bool) (*bst_t) {

	p.mutex.Lock()
	log.Printf("delete(%d)\n", key)


	var flag_found bool = false
	for {

		rc := p.deleteInner(key)
		if (rc == nil) {

			break
		}


		flag_found = true


		if (flag_uniq == true) {

			break
		}
	}


	if (flag_found == false) {

		p.mutex.Unlock()
		return nil
	}


	p.mutex.Unlock()
	return p
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func searchHandler(bst *bst_t, flag_uniq bool, w http.ResponseWriter, r *http.Request) {

	if (r.Method != "GET") {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":true, \"err_msg\":\"invalid method for /search\", \"is_found\":false }}\n")
		log.Printf("ERROR[search]: invalid method \"%s\"\n", r.Method)
		return
	}


	val := r.URL.Query()["val"][0]
	val_int64, err := strconv.ParseInt(val, 10, 64);
	if (err != nil) {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":true, \"err_msg\":\"invalid value for /search\", \"is_found\":false }}\n")
		log.Printf("ERROR[search]: invalid value \"%s\"\n", val)
		return
	}


	var p *bst_item_t
	p = bst.find(val_int64)
	if (p == nil) {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":false, \"err_msg\":\"\", \"is_found\":false }}\n")
		log.Printf("INFO[search]: value \"%s\" is not found\n", val)

	} else {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":false, \"err_msg\":\"\", \"is_found\":true }}\n")
		log.Printf("INFO[search]: value \"%s\" is found\n", val)
	}
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func insertHandler(bst *bst_t, flag_uniq bool, w http.ResponseWriter, r *http.Request) {

	if (r.Method != "POST") {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":true, \"err_msg\":\"invalid method for /insert\", \"is_inserted\":false }}\n")
		log.Printf("ERROR[insert]: invalid method \"%s\"\n", r.Method)
		return
	}


	err := r.ParseForm();
	if (err != nil) {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":true, \"err_msg\":\"invalid value for /insert\", \"is_inserted\":false }}\n")
		log.Printf("ERROR[insert]: can not parse data\n")
		return
	}


	val := r.FormValue("val")
	val_int64, err := strconv.ParseInt(val, 10, 64);
	if (err != nil) {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":true, \"err_msg\":\"invalid value for /insert\", \"is_inserted\":false }}\n")
		log.Printf("ERROR[insert]: invalid value \"%s\"\n", val)
		return
	}


	var p *bst_item_t
	p = bst.insert(val_int64, flag_uniq)
	if (p == nil) {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":false, \"err_msg\":\"\", \"is_inserted\":false }}\n")
		log.Printf("INFO[insert]: value \"%s\" is already exist\n", val)

	} else {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":false, \"err_msg\":\"\", \"is_inserted\":true }}\n")
		log.Printf("INFO[insert]: value \"%s\" is inserted\n", val)
	}
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
func deleteHandler(bst *bst_t, flag_uniq bool, w http.ResponseWriter, r *http.Request) {

	if (r.Method != "DELETE") {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":true, \"err_msg\":\"invalid method for /delete\", \"is_deleted\":false }}\n")
		log.Printf("ERROR[delete]: invalid method \"%s\"\n", r.Method)
		return
	}


	val := r.URL.Query()["val"][0]
	val_int64, err := strconv.ParseInt(val, 10, 64);
	if (err != nil) {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":true, \"err_msg\":\"invalid value for /delete\", \"is_deleted\":false }}\n")
		log.Printf("ERROR[delete]: invalid value \"%s\"\n", val)
		return
	}


	var p *bst_t
	p = bst.delete(val_int64, flag_uniq)
	if (p == nil) {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":false, \"err_msg\":\"\", \"is_deleted\":false }}\n")
		log.Printf("INFO[delete]: value \"%s\" is not found\n", val)

	} else {

		fmt.Fprintf(w, "{\"result\":{\"is_error\":false, \"err_msg\":\"\", \"is_deleted\":true }}\n")
		log.Printf("INFO[delete]: value \"%s\" is deleted\n", val)
	}
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



	http.HandleFunc("/search/", func(w http.ResponseWriter, r *http.Request) {

		searchHandler(&bst, flag_uniq, w, r)
	})
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {

		insertHandler(&bst, flag_uniq, w, r)
	})
	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {

		deleteHandler(&bst, flag_uniq, w, r)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "for search use: curl -X 'GET' -L 'http://localhost:8080/search?val=777'\n")
		fmt.Fprintf(w, "for insert use: curl -X 'POST' --data 'val=777' -L 'http://localhost:8080/insert'\n")
		fmt.Fprintf(w, "for delete use: curl -X 'DELETE' -L 'http://localhost:8080/delete?val=777'\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
/* ***************************************************************************************************************************************************************************************************************************************************************************************************************** */
