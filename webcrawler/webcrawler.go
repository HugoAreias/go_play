package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func Child(url string, chn chan string) {
    defer close(chn)

    body, urls, err := fetcher.Fetch(url)
    // fmt.Printf("Child: %s\n", url)
	if err != nil {
        fmt.Println(err)
		return
	}

    fmt.Printf("found: %s %q\n", url, body)

    for _, u := range urls {
        // fmt.Printf("sending %s\n", u)
        chn <- u
    }
}
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, fetcher Fetcher, seen seenUrl) {
    chn := make(chan string)
    go Child(url, chn)
    seen[url] = 1

    for u := range chn {
        // fmt.Printf("got %s\n", u)
        if _, ok := seen[u]; !ok {
            // fmt.Printf("go %s\n", u)
            Crawl(u, fetcher, seen)
        }
    }
}

func main() {
	Crawl("http://golang.org/", fetcher, seenUrl{})
}

type seenUrl map[string]int

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
