package taskrunner

const (
	// controlChan 里面的消息
	READY_TO_DISPATCH = "d" // 任务准备好了，可以分发了
	READY_TO_EXECUTE  = "e" // 任务准备好了，可以执行了
	CLOSE             = "c" // 任务关闭

	VIDEO_PATH = "goStreaming-on-demand-services/videos"
)

// 架构图里面的控制通道 chan string，只能传输字符串
type controlChan chan string

// 下发的数据通道，可以传输任意类型的数据
type dataChan chan interface{}

// dispatcher和executor的函数类型
type fn func(dc dataChan) error
