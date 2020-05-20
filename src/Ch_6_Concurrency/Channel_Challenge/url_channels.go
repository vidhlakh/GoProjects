// Get content type of sites
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("%s -> error: %s\n", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s\n", url, ctype)
}

func main() {

	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}
	//Create response channel
	ch := make(chan string)
	for _, url := range urls {

		go returnType(url, ch)
	}

	for range urls { // run number of urls
		out := <-ch
		fmt.Printf("Receiving url:%s", out)
	}

}
