package main

import (
	"fmt"
	"time"
)

func processTask(taskID int, ch chan string) {
	fmt.Printf("Task %d started\n", taskID)

	time.Sleep(time.Second)

	message := fmt.Sprintf("Task %d completed", taskID)
	ch <- message
}

func main() {
	fmt.Println("Starting Task Processor...")

	ch := make(chan string)

	for i := 1; i <= 3; i++ {
		go processTask(i, ch)
	}

	for i := 1; i <= 3; i++ {
		result := <-ch
		fmt.Println(result)
	}

	fmt.Println("All tasks completed!")
}
