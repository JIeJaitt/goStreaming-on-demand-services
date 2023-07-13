package defs

// requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// response
type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

// Comment
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

// Session
type SimpleSession struct {
	Username string // 登录名
	TTL      int64  // TTL 是 "Time To Live" 的缩写，表示 session 的存活时间。
}
