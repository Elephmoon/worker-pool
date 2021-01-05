package main

import (
	"fmt"
	"github.com/Elephmoon/worker-pool/workerpool"
	"math/rand"
	"time"
)

func main() {
	var tasks []*workerpool.Task
	pool := workerpool.NewPool(tasks, 5)
	go func() {
		for {
			taskID := rand.Intn(100) + 10
			if taskID%9 == 0 {
				pool.StopBackground()
			}
			time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
			task := workerpool.NewTask(func(data interface{}) error {
				id := data.(int)
				time.Sleep(time.Millisecond)
				fmt.Println("Task processed: ", id)
				return nil
			}, taskID)
			pool.AddTask(task)
		}
	}()
	pool.RunBackground()
}
