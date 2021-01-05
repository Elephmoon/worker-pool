package workerpool

import (
	"sync"
	"time"
)

type Pool struct {
	Tasks         []*Task
	Workers       []*Worker
	runBackground chan struct{}
	concurrency   int
	collector     chan *Task
	wg            sync.WaitGroup
}

func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		collector:   make(chan *Task, len(tasks)),
		wg:          sync.WaitGroup{},
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.concurrency; i++ {
		worker := NewWorker(p.collector, i)
		worker.Start(&p.wg)
	}
	for _, task := range p.Tasks {
		p.collector <- task
	}
	close(p.collector)
	p.wg.Wait()
}

func (p *Pool) AddTask(task *Task) {
	p.collector <- task
}

func (p *Pool) RunBackground() {
	go func() {
		time.Sleep(time.Millisecond)
	}()
	for i := 0; i < p.concurrency; i++ {
		worker := NewWorker(p.collector, i)
		p.Workers = append(p.Workers, worker)
		go worker.StartBackground()
	}
	for _, task := range p.Tasks {
		p.collector <- task
	}
	p.runBackground = make(chan struct{})
	<-p.runBackground
}

func (p *Pool) StopBackground() {
	for _, worker := range p.Workers {
		worker.StopBackground()
	}
	p.runBackground <- struct{}{}
}
