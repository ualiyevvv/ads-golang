package goroutines

import (
	"fmt"
	"sync"
)

func a() {
	fmt.Println("a")
	ch1 <- struct{}{}
	defer wg.Done()
}

func b() {
	<-ch1
	fmt.Println("b")
	ch2 <- struct{}{}
	defer wg.Done()
}

func c() {
	<-ch2
	fmt.Println("c")
	defer wg.Done()
}

var (
	ch1 = make(chan struct{})
	ch2 = make(chan struct{})
)
var wg sync.WaitGroup

func Start() {
	wg.Add(3)
	go a()
	go b()
	go c()
	wg.Wait()

}
