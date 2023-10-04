package main

import (
	"fmt"
	"time"

	"github.com/Jeenusha1/Golang/channels"
	
)
// go routines[managed by go runtime] - concurrency vs parallelism
// go keyword - if used without proper handling of thread, may skip the task
func main() {
	go prints("hi \n")
	prints("jee \n")
	// how channels work
	channels.Chann()
	///

}

func prints(s string) {
	for i := 0; i < 3; i++ {
		// on use of go routines, should give time for executing thread
		// sync package can be used [provide sychronization]
		time.Sleep(3 * time.Millisecond)
		fmt.Print(s)

	}
}


