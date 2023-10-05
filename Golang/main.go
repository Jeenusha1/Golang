package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Jeenusha1/Golang/async"
	"github.com/Jeenusha1/Golang/channels"
	"github.com/Jeenusha1/Golang/contextss"
)

// go routines[managed by go runtime] - concurrency vs parallelism
// go keyword - if used without proper handling of thread, may skip the task
func main() {
	go prints("hi \n")
	prints("jee \n")
	// how channels work
	channels.Chann()
	//synchronization - channel helps us to execute the goroutine one by one
	// informs the other process that other goroutine is done
	done := make(chan bool, 1)
	go channels.Check(done)
	// done here inform where the other goroutine is completed
	// without <-done (channel receives the info) check2 can be executed first
	if <-done {
		go channels.Check2()
		time.Sleep(time.Second)
	}

	//pushing the slice to function with go keyword to execute asynchr.
	//without go keyword the urls executes in order
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	for _, url := range urls {
		//if we dont use the go keyword the function executes the url one by one (synchr.)
		go async.CheckUrl(url)
	}
	// on use of go keyword, before goroutine finishes main routine will be closed
	// using sleep keeps the main to wait till the goroutine finishes
	// using sleep is not reliable
	time.Sleep(5 * time.Second)

	//using channel to make the main and called function to communicate
	channels.Callforuseasync()

	/// how context work
	// Background has no value, never terminates, empty context (non nil)
	ctx := context.Background()
	// empty context [gets the key value on below function]
	ctx = contextss.GetContext(ctx)
	// accessed key value by using ctx
	val := ctx.Value("jee")
	fmt.Print(val)

}

func prints(s string) {
	for i := 0; i < 3; i++ {
		// on use of go routines, should give time for executing thread
		// sync package can be used [provide sychronization]
		time.Sleep(3 * time.Millisecond)
		fmt.Print(s)

	}
}
