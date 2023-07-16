package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	// 通过 sc(status code) 设置状态码
	w.WriteHeader(sc)
	// 通过 io.WriteString 写入错误信息返回给客户端
	io.WriteString(w, errMsg)
}
