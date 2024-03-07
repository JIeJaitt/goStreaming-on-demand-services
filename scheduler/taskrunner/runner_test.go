package taskrunner

import (
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %v", i)
		}
		return nil
	}

	e := func(dc dataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("EXECUTE recevied: %v", d)

			default:
				break forloop
			}
		}
		//return errors.New("exit")
		return nil
	}

	runner := NewRunner(30, false, d, e)
	// 不加goroutine会一直循环下去
	//runner.StartAll()
	// 加了goroutine才会执行到下一行
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}
