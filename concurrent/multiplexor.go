package concurrent

import "sync"

func Funnel(in []chan int) <-chan int {
	dest := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(in))

	for _, c := range in {
		go func(c <-chan int) {
			defer wg.Done()
			for msg := range c {
				dest <- msg
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest
}
