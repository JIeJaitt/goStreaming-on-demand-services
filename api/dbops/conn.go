package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

// init 函数会在 main 函数之前执行
// 也就是说 init 函数会在程序加载的时候第一个被执行
func init() {
	// 如果数据库都没连接起来, 那么服务端跑起来也没什么意义
	// 在项目中, 不要轻易使用 panic
	// panic 会中断整个项目, 并且抛出这个 panic
	// 只有在确实无法正常完成业务逻辑的时候, 才使用 panic
	dbConn, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/stmsrv?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

}
