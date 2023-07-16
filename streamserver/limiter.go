package main

import "log"

// ConnLimiter 是一个结构体，用于限制并发连接数
type ConnLimiter struct {
	concurrentConn int      // 并发连接数
	bucket         chan int // 令牌桶通道
}

// NewConnLimiter 是 ConnLimiter 的构造函数
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc), // 缓冲区大小为 cc
	}
}

func (cl *ConnLimiter) GetConn() bool {
	// 如果通道已满，说明已经达到并发连接数的上限，返回 false
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation.")
		return false
	}
	// 如果通道未满，说明还有空闲的连接，返回 true
	cl.bucket <- 1
	return true
}

// ReleaseConn 用于释放连接
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connection coming: %d", c)
}
