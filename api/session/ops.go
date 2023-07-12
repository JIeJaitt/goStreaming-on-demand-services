package session

import (
	"goStreaming-on-demand-services/api/dbops"
	"goStreaming-on-demand-services/api/defs"
	"goStreaming-on-demand-services/api/utils"
	"sync" // 储存 session 信息
	"time" // 判断 session 是否过期
)

var sessionMap *sync.Map

func init() {

	sessionMap = &sync.Map{}
}

func noewInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// LoadSessionsFromDB 从 db 中加载到 map(cache) 里面
func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

// GenerateNewSessionId 生成一个新的 session
func GenerateNewSessionId(username string) string {
	id, _ := utils.NewUUID()
	ct := noewInMilli()
	ttl := ct + 30*60*1000 // Severside session valid time: 30 min

	ss := &defs.SimpleSession{
		Username: username,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)
	dbops.InserSession(id, ttl, username)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := noewInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			// delete expired session
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}

	return "", true
}
