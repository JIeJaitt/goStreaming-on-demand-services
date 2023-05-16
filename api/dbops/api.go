package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"stmsrv/api/defs"
	"stmsrv/api/utils"
	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		// 如果我们在这里直接返回 err
		// 那么后面的 stmtIns.Close() 语句就不会执行
		// 所以后面还是需要用 defer 来执行Close
		return err
	}

	// defer 是在栈退出的时候才会调用
	// 所以 defer 对性能有些许损耗
	// 尽量少在对性能要求严格的项目中使用 defer
	//
	// 在整个函数中，跳出的机会非常多
	// 如果不使用 defer，就没法确保在函数退出时关闭 stmtIns
	// 除非在每次跳出的时候都写一遍 stmtIns.Close()
	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			return
		}
	}(stmtIns)
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		// string 是没有 nil 的, 所以这里返回一个空字符串
		// nil 大多数时候是出现在 pointer, interface, map, slice, channel, function 这些类型上
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	// ErrNoRows 实际上并不是一个真正的处理错误
	// 而是没有结果
	// 而这里如果不处理程序会把 NoRows 按照错误来返回
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer func(stmtOut *sql.Stmt) {
		err := stmtOut.Close()
		if err != nil {
			return
		}
	}(stmtOut)
	return pwd, nil
}

func DeleteUserCredential(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer func(stmtDel *sql.Stmt) {
		err := stmtDel.Close()
		if err != nil {
			return
		}
	}(stmtDel)
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	// create time -> db ->
	// 真正的写库时间和生成视频的时间是有区别的

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info 
		(id, author_id, name, display_ctime) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	defer func(stmtIns *sql.Stmt) {
		err := stmtIns.Close()
		if err != nil {
			return
		}
	}(stmtIns)

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return res, nil
}

type UserError struct {
	Message string
}

func (e *UserError) Error() string {
	return e.Message
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT author_id, name, display_ctime FROM video_info 
		WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	var aid int
	var dct string
	var name string
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)

	if err == sql.ErrNoRows {
		return nil, &UserError{"User not found"}
	}
	if err != nil {
		return nil, &UserError{"Error retrieving user"}
	}
	defer func(stmtOut *sql.Stmt) {
		err := stmtOut.Close()
		if err != nil {
			return
		}
	}(stmtOut)

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer func(stmtDel *sql.Stmt) {
		err := stmtDel.Close()
		if err != nil {
			return
		}
	}(stmtDel)
	return nil
}
