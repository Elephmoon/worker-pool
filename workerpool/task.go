package workerpool

type Task struct {
	Err  error
	Data interface{}
	fn   func(interface{}) error
}

func NewTask(fn func(interface{}) error, data interface{}) *Task {
	return &Task{
		Data: data,
		fn:   fn,
	}
}

func process(task *Task) {
	task.Err = task.fn(task.Data)
}
