package taskrunner

import "log"

type Runner struct {
	Controller controlChan // 控制通道
	Error      controlChan // 错误通道
	Data       dataChan    // 数据通道
	dataSize   int         // 数据通道的大小
	longLived  bool        // 是否长期存活
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longLived:  longlived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

// 常驻任务
func (r *Runner) startDispatch() {
	// 非常驻函数杀掉进程
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for {
		select {
		case c := <-r.Controller:
			// 处理状态为DISPATCH情况
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				log.Printf("startDispatch Controller add data: %v\n", r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					// 改变Controller状态
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				log.Printf("startDispatch Controller execute data: %v\n", r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}

		// 处理出错情况
		case e := <-r.Error:
			if e == CLOSE {
				return
			}
			//default:

		}
	}
}

// 启动ruuner
func (r *Runner) StartAll() {
	// 启动前需要前内置一个READY_TO_DISPATCH来激活程序
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
