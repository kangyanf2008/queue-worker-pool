package worker

import "fmt"

type Worker struct {
	JobQueue chan Job
	Quit     chan bool

}

//创建对象
func NewWorker() *Worker  {
	return &Worker{JobQueue: make(chan Job)}
}

//运行工作线程
func (w Worker) Run(workerQueue chan chan Job)  {
	go func() {
		for {
			//注册工作线程
			workerQueue <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				if err := job.Do(); err != nil {
					fmt.Printf("execute Job error:%v \n", err)
				}
			case <-w.Quit:
				close(w.JobQueue)
				return
			}
		}
	}()
}