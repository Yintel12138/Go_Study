package main

import (
	"fmt"
	"runtime"
	"sync"
)

//func main() {
//	fmt.Println("hello world!")
//}
func testRoutine(i int) {
	fmt.Println(i)
}

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("counter = ", counter)
	lock.Unlock()
}
func main() {
	mutex := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(&mutex)
		//fmt.Println(i)
	}
	for true {
		mutex.Lock()
		c := counter
		mutex.Unlock()

		runtime.Gosched()

		if c >= 10 {
			break
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
