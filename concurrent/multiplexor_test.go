package concurrent

import (
	"sync"
	"testing"
)

func TestFunnel(t *testing.T) {
	N := 5
	ch1 := make(chan int)
	ch2 := make(chan int)

	res := Funnel([]chan int{ch1, ch2})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		wg.Done()
		for i := 0; i < N; i++ {
			ch1 <- i
			ch2 <- i
		}
		close(ch1)
		close(ch2)
	}()

	wg.Wait()

	for r := range res {
		println(r)
	}
}
