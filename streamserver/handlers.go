package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 从 URL 中获取 vid-id
	vid := p.ByName("vid-id")
	// 获取文件目录
	vl := VIDEO_DIR + vid

	// 打开文件
	video, err := os.Open(vl)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error.")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	// 把文件作为二进制流传输给客户端
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
