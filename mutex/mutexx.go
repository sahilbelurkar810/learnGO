package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(w *sync.WaitGroup) {
	p.mu.Lock()
	p.views += 1
	defer func() {
		w.Done()
		p.mu.Unlock()
	}()

}

func main() {

	wg := sync.WaitGroup{}
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}

	wg.Wait()

	fmt.Println("Post views:", myPost.views)
}
