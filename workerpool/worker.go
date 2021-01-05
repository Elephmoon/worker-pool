package workerpool

import "sync"

type Worker struct {
	ID         int
	taskChan   chan *Task
	finishChan chan struct{}
}

func NewWorker(taskChan chan *Task, ID int) *Worker {
	return &Worker{
		ID:         ID,
		taskChan:   taskChan,
		finishChan: make(chan struct{}),
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

func (w *Worker) StartBackground() {
	for {
		select {
		case task := <-w.taskChan:
			process(task)
		case <-w.finishChan:
			return
		}
	}
}

func (w *Worker) StopBackground() {
	w.finishChan <- struct{}{}
}
