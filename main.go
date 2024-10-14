package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// If a channel isn't used, the main routine would just execute until the end
		// as it won't wait for the child go routines to finish
		go checkLink(link, c)
	}

	// If a loop isn't used, the print would just execute once since it's only waiting for one message from the channel
	// then it exits
	// Using the range keyword, it's telling the for loop to wait for a value in the c channel, before assigning it to l
	// then passing that value into checkLink
	for l := range c {
		// Originally the value outside the scope is supposed to be passed into the function literal so that the value can be used instead of the reference
		// like so:
		// go func(link string) {
		// 	time.Sleep(5 * time.Second)
		// 	// Reading from the channel is a blocking call.
		// 	checkLink(link, c)
		// }(l)
		// But as of https://go.dev/blog/go1.22 there is no need to do so anymore
		go func() {
			time.Sleep(5 * time.Second)
			// Reading from the channel is a blocking call.
			checkLink(l, c)
		}()
	}

}

func checkLink(link string, c chan string) {
	// All Go functions are by default "synchronous" in Javascript terms
	// So this function will block all other processes until it is done
	// Go considers this a blocking call
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
