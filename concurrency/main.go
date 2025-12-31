package main

import (
	"fmt"
	"time"
)

func main() {
	dones := make([]chan bool, 4)

	for i := 0; i < 4; i++ {
		dones[i] = make(chan bool)
	}
	go slowGreet("test slow", dones[0])
	go greet("Test", dones[1])
	go slowGreet("test slow2", dones[2])
	go greet("Test2", dones[3])

	for _, done := range dones {
		<-done
	}

	singleChan := make(chan bool)
	go greet("Single chan greet", singleChan)
	go slowGreet("Single chan greet", singleChan)

	for range singleChan {

	}
}

func greet(content string, doneChan chan bool) {
	fmt.Println(content)
	doneChan <- true
	close(doneChan)
}

func slowGreet(content string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(content)

	doneChan <- true
	close(doneChan)
}
