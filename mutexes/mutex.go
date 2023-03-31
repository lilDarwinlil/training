package mutexes

import (
	"fmt"
	"sync"
)

var counter int

func sum(wg *sync.WaitGroup, m *sync.Mutex) {
	//time.Sleep(time.Millisecond * 10)
	//m.Lock()
	counter = counter + 1
	//m.Unlock()
	wg.Done()
}

func RunMutex() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go sum(&wg, &m)
	}

	wg.Wait()

	fmt.Println("value of counter after 1000 operations is", counter)
}
