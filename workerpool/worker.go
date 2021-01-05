package workerpool

import "sync"

type Worker struct {
	ID       int
	taskChan chan *Task
}

func NewWorker(taskChan chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: taskChan,
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for task := range w.taskChan {
			process(task)
		}
	}()
}
