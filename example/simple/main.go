package main

import (
	"fmt"
	"github.com/Elephmoon/worker-pool/workerpool"
	"time"
)

func main() {
	var tasks []*workerpool.Task
	for i := 0; i < 1000; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Task processed: ", taskID)
			return nil
		}, i)
		tasks = append(tasks, task)
	}
	pool := workerpool.NewPool(tasks, 5)
	pool.Run()
}
