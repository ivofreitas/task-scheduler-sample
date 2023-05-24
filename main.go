package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Task struct {
	ID   int
	Data interface{}
}

type Pool struct {
	TaskCh   chan Task
	ResultCh chan Task
	Size     int
}

func (p *Pool) Worker(id int) {
	for task := range p.TaskCh {
		// Perform task
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		time.Sleep(time.Duration(n) * time.Millisecond)
		task.Data = fmt.Sprintf("Task #%v completed by Worker #%v", task.ID, id)

		p.ResultCh <- task
	}
}

func TaskScheduler(tasks []Task, numWorkers int) []Task {
	// Create channels for task and result communication
	taskChannel := make(chan Task, len(tasks))
	resultChannel := make(chan Task, len(tasks))

	workerPool := Pool{
		TaskCh:   taskChannel,
		ResultCh: resultChannel,
		Size:     numWorkers,
	}
	// Start the worker pool
	for i := 0; i < numWorkers; i++ {
		go workerPool.Worker(i)
	}

	// Distribute tasks among workers
	go func() {
		for _, task := range tasks {
			taskChannel <- task
		}
		close(taskChannel)
	}()

	// Collect the processed tasks from the result channel
	var results []Task
	for i := 0; i < len(tasks); i++ {
		result := <-resultChannel
		results = append(results, result)
	}

	close(resultChannel)

	return results
}

func main() {
	tasks := []Task{
		{ID: 1, Data: "Task 1 data"},
		{ID: 2, Data: "Task 2 data"},
		{ID: 3, Data: "Task 3 data"},
		{ID: 4, Data: "Task 4 data"},
		{ID: 5, Data: "Task 5 data"},
		{ID: 6, Data: "Task 6 data"},
		{ID: 7, Data: "Task 7 data"},
		{ID: 8, Data: "Task 8 data"},
		{ID: 9, Data: "Task 9 data"},
	}

	numWorkers := 3

	results := TaskScheduler(tasks, numWorkers)

	// Print the results
	for _, result := range results {
		fmt.Println(result.Data)
	}
}
