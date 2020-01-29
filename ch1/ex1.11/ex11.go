// ex11 fetches URLs in parallel and reports their times and sizes.
// The alexa top million web sites can be downloaded from
// http://s3.amazonaws.com/alexa-static/top-1m.csv.zip
// Here is a shortcut for running a given set of URLs.
// Assuming you extracted the top-1m.csv.zip in the same
// directory as the code, you can append any number of URLs
// by typing the following command:
// go run ex11.go $(cat top-1m.csv | head -20 | awk -F, '{print$2}' | sed 's/^/https:\/\//'| tr '\n' ' ')
// The above will append the first 20 URLs from the file. You can use 'tail' instead of head to start
// from the bottom
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
