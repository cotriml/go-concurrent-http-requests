package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Text string
}

func main() {
	usualHTTPCalls()
	concurrentHTTPCalls()
}

func concurrentHTTPCalls() {
	start := time.Now()
	defer printExecutionTime(start)

	ch := make(chan Message)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			message := doHTTPCall(i)
			ch <- message
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for message := range ch {
		fmt.Printf("Fetched message %v\n", message.Text)
	}

}

func usualHTTPCalls() {
	start := time.Now()
	defer printExecutionTime(start)

	for i := 0; i < 20; i++ {
		message := doHTTPCall(i)
		fmt.Printf("Fetched message %v\n", message.Text)
	}
}

func doHTTPCall(index int) Message {
	time.Sleep(1 * time.Second)
	return Message{Text: fmt.Sprintf("Message %v", index+1)}
}

func printExecutionTime(start time.Time) {
	fmt.Println("Execution Time: ", time.Since(start))
}
