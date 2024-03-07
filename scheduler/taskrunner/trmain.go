package taskrunner

import (
	"log"
	"time"
)

// Timer部分
// setup -> start { trigger -> task -> runner } -> close

type Worker struct {
	ticket *time.Ticker // 定时器
	runner *Runner      // 任务执行器
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticket: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		// 定时器到了
		case <-w.ticket.C:
			log.Printf("ticket run start--------------------\n")
			go w.runner.StartAll()
			log.Printf("ticket run end--------------------\n")
		}
	}
}

func Start() {
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}
