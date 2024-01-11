package worker

//工作协和池
type WokerPool struct {
	Worker []*Worker
	quit chan bool
	WorkerQueue chan chan Job
	queue chan Job
}

//实例化协程池
func NewWokerPool(wokerNum int) *WokerPool {
	maxWokerNum := wokerNum
	if maxWokerNum <= 0 {
		maxWokerNum = MAX_QUEUE_SIZE
	}
	return &WokerPool{
			WorkerQueue: make(chan chan Job, Max_Worker_Pool_Size), //工作协程数
			queue: make(chan Job, maxWokerNum)}                  //请求接收队列数
}

//任务放入任务队列
func(wp *WokerPool) PutJob(j Job){
	wp.queue <- j
}

//启动工作协和
func (wp *WokerPool) Run() *WokerPool {
	for i:=0; i < Max_Worker_Pool_Size; i++ {
		worker := NewWorker()
		wp.Worker = append(wp.Worker, worker)
		worker.Run(wp.WorkerQueue)   //工作协和队列
	}
	go func() {
		for {
			select {
			case job := <-wp.queue:
				jobChan := <- wp.WorkerQueue
				jobChan<-job
			case <-wp.quit:
				//工作线程结束处理
				for _, w:= range wp.Worker {
					w.Quit<-true
				}
				return
			}
		}
	}()
	return wp
}

func (d *WokerPool) Stop(){
	d.quit<-true
}
