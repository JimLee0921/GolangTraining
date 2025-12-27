package main

import "sync"

type httpPkg struct {
}

func (h httpPkg) Get(url string) {

}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"https://go.dev/",
		"https://pkg.go.dev/",
	}

	for _, url := range urls {
		wg.Go(func() {
			http.Get(url)
		})
	}
	// 等待全部 HTTP 获取完毕
	wg.Wait()
}
