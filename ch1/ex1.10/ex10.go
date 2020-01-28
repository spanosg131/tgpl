// ex10 fetches URLs in parallel and reports their times and sizes.
// Running ex10 multiple times on a single page did not improve the
// load time which means no caching is taking place.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

func main() {
	f, err := os.Create("results.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		f.WriteString(<-ch)
		//fmt.Println(<-ch) // receive from channel
	}

	totalElapsed := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	f.WriteString(totalElapsed)
}
