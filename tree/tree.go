package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
    "time"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        Walk(t.Left, ch)
    }
	ch <- t.Value
    if t.Right != nil {
    	Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1, ch2 := make(chan int), make(chan int) 
	go Walk(t1, ch1)
    go Walk(t2, ch2)

    for {
        select {
        case val1 := <-ch1:
            val2 := <-ch2
            if val1 != val2 {
            	return false
        	}
        case <-time.After(500 * time.Millisecond):
            return true
        }
    }
}

func main() {
	//ch := make(chan int)
	//go Walk(tree.New(1), ch)
	//for i := 0; i < 10; i++ {
		//fmt.Println(<-ch)
	//}
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(1), tree.New(2)))
}
