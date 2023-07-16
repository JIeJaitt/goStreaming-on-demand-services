package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("../videos/upload.html")
	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 从 URL 中获取 vid-id
	vid := p.ByName("vid-id")
	// 获取文件目录
	vl := VIDEO_DIR + vid

	// 打开文件
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error.")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	// 把文件作为二进制流传输给客户端
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

// uploadHandler是一个处理上传文件请求的处理函数
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 设置请求体的最大字节数限制为MAX_UPLOAD_SIZE
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	// 解析多部分表单数据，其中包含上传的文件
	err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
	if err != nil {
		// 发生错误时，向客户端发送400错误响应
		sendErrorResponse(w, http.StatusBadRequest, "Failed to parse multipart form")
		return
	}

	// 从表单中获取上传的文件
	file, _, err := r.FormFile("file")
	if err != nil {
		// 发生错误时，向客户端发送400错误响应
		sendErrorResponse(w, http.StatusBadRequest, "Failed to get file from form")
		return
	}
	defer file.Close()

	// 读取文件的内容
	data, err := io.ReadAll(file)
	if err != nil {
		// 发生错误时，记录错误日志并向客户端发送500错误响应
		log.Printf("Failed to read file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to read file")
		return
	}

	// 获取URL中的vid-id参数
	vid := p.ByName("vid-id")

	// 构建文件的完整路径
	filePath := VIDEO_DIR + vid

	// 将文件内容写入到指定路径的文件中
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		// 发生错误时，记录错误日志并向客户端发送500错误响应
		log.Printf("Failed to write file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Failed to write file")
		return
	}

	// 向客户端发送201状态码，表示文件上传成功
	w.WriteHeader(http.StatusCreated)
	// 向客户端发送成功消息
	io.WriteString(w, "File uploaded successfully")
}
