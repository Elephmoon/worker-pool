package workerpool

import "sync"

type Pool struct {
	Tasks       []*Task
	concurrency int
	collector   chan *Task
	wg          sync.WaitGroup
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
