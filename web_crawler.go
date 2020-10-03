package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var crawled map[string]bool

func childCrawl(url string, depth int, fetcher Fetcher, result chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		return
	}

	if e := crawled[url]; e {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		result <- err.Error()
		return
	}

	crawled[url] = true

	result <- fmt.Sprintf("found: %s %q", url, body)
	for _, u := range urls {
		wg.Add(1)
		go childCrawl(u, depth-1, fetcher, result, wg)
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	result := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go childCrawl(url, depth, fetcher, result, &wg)

	go func() {
		wg.Wait()
		close(result)
	}()
	for r := range result {
		fmt.Println(r)
	}
}

func main() {
	crawled = make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
