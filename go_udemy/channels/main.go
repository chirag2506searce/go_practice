package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://golang.org",
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://flipkart.com",
		"http://apple.com",
		"http://netflix.com",
	}
	// fmt.Println(links)
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// do above or {
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// }
	for l := range c {
		// for {
		// go checkLink(<-c, c)

		// go checkLink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}
	// fmt.Println(<-c) // will hang, because len(links) = 8

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		// c <- "Its down"
		c <- link
		return
	}
	fmt.Println(link, "is up")
	// c <- "Its up"
	c <- link
}
