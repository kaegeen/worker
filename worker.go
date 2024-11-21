package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	id    int
	value int
}

func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		// Simulate processing time
		time.Sleep(time.Millisecond * time.Duration(task.value%100))
		result := task.value * task.value
		fmt.Printf("Worker %d processed task %d: %d^2 = %d\n", id, task.id, task.value, result)
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	// Channel to send tasks to workers
	tasks := make(chan Task, numTasks)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Generate tasks and send them to the channel
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{id: i, value: i * 10} // Each task is a number that gets squared
	}

	// Close the task channel after all tasks are pushed
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
}
