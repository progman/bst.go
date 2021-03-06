package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"sync"
	"fmt"
	"net/http"
	"strconv"
	"os"
)


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  BST item (leaf)
 */
type bst_item_t struct {

	parent *bst_item_t
	left   *bst_item_t
	right  *bst_item_t
	key int64
}


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  BST item init
 * @param  key integer value
 * @return void
 */
func (p *bst_item_t) init(key int64) {

	p.parent = nil
	p.left   = nil
	p.right  = nil
	p.key    = key
}


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  BST main struct
 */
type bst_t struct {

	head       *bst_item_t
	mutex      sync.RWMutex
	debug_list []*bst_item_t
	flag_debug bool
}


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  BST init
 * @return void
 */
func (p *bst_t) init(flag_debug bool) {

	p.head       = nil
	p.flag_debug = flag_debug
}


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  insert key to BST (we do not use recursion becase we think about stack overflow)
 * @param  key integer value
 * @param  flag_uniq flag for uniq values in BST
 * @return pointer to BST item (leaf)
 */
func (p *bst_t) insert(key int64, flag_uniq bool) (*bst_item_t) {

	p.mutex.Lock()
	log.Printf("insert(%d)\n", key)
	var p_bstr_item *bst_item_t = &bst_item_t{}

	if (p.flag_debug == true) {
		p.debug_list = append(p.debug_list, p_bstr_item)
	}


	p_bstr_item.init(key)


	if (p.head == nil) {

		log.Printf("\tset as root\n")
		p.head = p_bstr_item
		p.mutex.Unlock()
		return p_bstr_item
	}


// we do not use recursion becase we think about stack overflow
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


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  find key in BST (we do not use recursion becase we think about stack overflow)
 * @param  key integer value
 * @return pointer to BST item (leaf)
 */
func (p *bst_t) findInner(key int64) (*bst_item_t) {

	if (p.head == nil) {

		log.Printf("\tbst is empty\n")
		return nil
	}


// we do not use recursion becase we think about stack overflow
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


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  find key in BST (we do not use recursion becase we think about stack overflow)
 * @param  key integer value
 * @return pointer to BST item (leaf)
 */
func (p *bst_t) find(key int64) (*bst_item_t) {

	p.mutex.RLock()
	log.Printf("find(%d)\n", key)


	p_cur := p.findInner(key)


	p.mutex.RUnlock()
	return p_cur
}


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  delete key in BST (we do not use recursion becase we think about stack overflow)
 * @param  key integer value
 * @return pointer to BST
 */
func (p *bst_t) deleteInner(key int64) (*bst_t) {

	var p_cur *bst_item_t = p.findInner(key)
	if (p_cur == nil) {

		return nil
	}


// we do not use recursion becase we think about stack overflow


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


// UNDER CONSTRUCTION !!!


	return p
}


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  delete key in BST (we do not use recursion becase we think about stack overflow)
 * @param  key integer value
 * @param  flag_uniq flag for uniq values in BST
 * @return pointer to BST
 */
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


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  http handler for operation search
 * @param  bst pointer to BST
 * @param  flag_uniq flag for uniq values in BST
 * @param  w response context
 * @param  r request context
 * @return void
 */
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


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  http handler for operation insert
 * @param  bst pointer to BST
 * @param  flag_uniq flag for uniq values in BST
 * @param  w response context
 * @param  r request context
 * @return void
 */
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


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  http handler for operation delete
 * @param  bst pointer to BST
 * @param  flag_uniq flag for uniq values in BST
 * @param  w response context
 * @param  r request context
 * @return void
 */
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


/**
 * @author Alexey Potehin <gnuplanet@gmail.com>, http://www.gnuplanet.online/doc/cv
 * @brief  main function
 */
func main() {

	var err error
	var i int
	var int_list []int64
//	var p *bst_item_t
	var bst bst_t
	var flag_uniq bool = false


// load int list from json file
	intJson, err := ioutil.ReadFile("./bst.json")
	if err != nil {

		log.Fatal(err)
	}
//	intJson := `[ 100, 110, 90, 100, 110, 120, 105 ]`


// parse json
	err = json.Unmarshal([]byte(intJson), &int_list)
	if err != nil {

		log.Fatal(err)
	}
	fmt.Printf("int_list: %+v\n", int_list)


// init BST
	bst.init(false)


// load values from int list to BST
	for i=0; i < len(int_list); i++ {

		bst.insert(int_list[i], flag_uniq)
	}



/*
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
*/


/*
	log.Printf("head: %p\n", bst.head)
	for i=0; i < len(list); i++ {

		log.Printf("item:        %p\n", list[i])
		log.Printf("item.parent: %p\n", list[i].parent)
		log.Printf("item.left:   %p\n", list[i].left)
		log.Printf("item.right:  %p\n", list[i].right)
		log.Printf("item.key:    %d\n", list[i].key)

		log.Printf("\n")
	}
*/


// setup http server
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


	os.Exit(0)
}
