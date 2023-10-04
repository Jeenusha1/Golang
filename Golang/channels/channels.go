package channels

import "fmt"

func Chann(){

	val := make(chan int)
	// go keyword plays vital role here - chan needs goroutine to work properly
	go receiveandkeep(val, 2)
	go receiveandkeep(val, 4)

	v1, v2 := <-val, <-val
	fmt.Print(v1, v2)

}

func receiveandkeep(c chan int, value int) {
	c <- value * 5

}