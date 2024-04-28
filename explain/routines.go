// Go Routines

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}


// Channels

ch := make(chan int)

// channel communication that transport int values

func say(s string, ch chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		ch <- s
	}
}

func main() {
	ch := make(chan string)
	go say("Hello", ch)

	for {
		msg := <-ch
		fmt.Println(msg)
	}

	fmt.Println("Goodbye")
}

// Buffered channels

ch := make(chan int, 100)

// channel with capacity of 100

func main(){
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	ch <- 3 // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Range and close

v, ok := <-ch
// ok is false if channel is closed, and true if channel is not closed

for i := range ch {
	// do something
}

// close channel is not necessary, it will be closed automatically
// when the loop ends or the channel is closed

func say(s string, ch chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		ch <- s
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	go say("Hello", ch)

	for {
		msg, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received: ", msg)
	}

	fmt.Println("Goodbye")
}

// Select

func say(s string, ch chan string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		ch <- s
	}
	close(ch)
}

func main() {
	ch := make(chan string)
	go say("Hello", ch)

	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				break
			}
			fmt.Println("Received: ", msg)

		case <-time.After(1 * time.Second): // timeout in 1 second time
			fmt.Println("Timeout in 1 second")
			break


		default:
			fmt.Println("Nothing received")
		}
	}
	fmt.Println("Goodbye")
}

// TreeNode Traversals example

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Tree Traversal
// Verify if two trees are the same
// Concurrent Traversal of two trees is the same sequence in the tree traversal order

func isSameSequence(root1, root2 *TreeNode) bool {
	seq1 := make(map[int]bool)
	seq2 := make(map[int]bool)

	transverse(root1, seq1)
	transverse(root2, seq2)
	
	return equal(seq1, seq2)

}


func transverse(node *TreeNode, seq map[int]bool) {
	if node == nil {
		return
	}
	seq[node.Val] = true

	transverse(node.Left, seq)
	transverse(node.Right, seq)
}

func equal(seq1, seq2 map[int]bool) bool {
	if len(seq1) != len(seq2) {
		return false
	}
	for val := range seq1 {
		if !seq2[val] {
			return false
		}
	}
	return true
}

func main() {
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 2,
			}
		},
		Right : &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 5
			},
			Right: &TreeNode{
				Val: 13
			}
		}
	}

	root2 := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2
				}
			},
			Right: &TreeNode{
				Val: 5,
			}
		},
		Right : &TreeNode{
			Val: 13
		}
	}

	fmt.Println(isSameSequence(root1, root2)) // true // because they are the same sequence in the tree traversal order of root1 and root2 both are 3, 1, 1, 2, 8, 5, 13 in the same order	

}
