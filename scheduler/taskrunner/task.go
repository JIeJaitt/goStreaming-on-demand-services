package taskrunner

import (
	"errors"
	"goStreaming-on-demand-services/scheduler/dbops"
	"log"
	"os"
	"sync"
)

// 删除物理文件
func deleteVideo(vid string) error {
	err := os.Remove(VIDEO_PATH + vid)

	// 当错误不是文件不存在错误时表示文件没有删除，将错误返回
	if err != nil && !os.IsNotExist(err) {
		log.Printf("Deleting video error: %v\n", err)
		return err
	}
	return nil
}

// 获取要删除的videoId加载到dataChan中
func VideoClearDispatcher(dc dataChan) error {
	log.Printf("start VideoClearDispatcher!\n")
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispatcher error: %v\n", err)
		return err
	}

	// 没有获取到数据，返回
	if len(res) == 0 {
		log.Printf("VideoClearDispatcher dataChan is zero : %v\n", err)
		return errors.New("all tasks finished")
	}

	// 将取到的id写道dataChan中
	for _, id := range res {
		log.Printf("VideoClearDispatcher add vid: %s\n", id)
		dc <- id
	}
	log.Printf("VideoClearDispatcher Normal ")
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	log.Printf("start VideoClearExecutor!\n")
	// 定义一个装错误的map
	errMap := &sync.Map{}
	var err error

forloop:
	for {
		select {
		case vid := <-dc:
			log.Printf("VideoClearExecutor get vid: %s\n", vid)
			go func(id interface{}) {
				// 删除数据库前先把文件删掉
				if err := deleteVideo(id.(string)); err != nil {
					log.Printf("VideoClearExecutor deleteVideo error: %v\n", err)
					errMap.Store(id, err)
					return
				}

				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					log.Printf("VideoClearExecutor DelVideoDeltionRecord error: %v\n", err)
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			log.Printf("VideoClearExecutor go to forloop!\n")
			break forloop
		}
	}

	// 遍历errMap
	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})
	log.Printf("errMap: %v\n", err)
	return err
}
