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

