package channels

import (
	"fmt"
	"net/http"
)

//using the struct type for commincation in channels
type urlStatus struct {
	url    string
	status bool
}

func CheckUrlforChannel(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		// The website is down
		c <- urlStatus{url, false}
	} else {
		// The website is up
		c <- urlStatus{url, true}
	}
}

func Callforuseasync() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	c := make(chan urlStatus)
	for _, url := range urls {
		go CheckUrlforChannel(url, c)

	}
	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, "is up.")
		} else {
			fmt.Println(result[i].url, "is down !!")
		}
	}
}
