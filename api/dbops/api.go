package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func openConn() *sql.DB {
	dbConn, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/stmsrv?charset=utf8")
	if err != nil {
		// 如果数据库都没连接起来, 那么服务端跑起来也没什么意义
		// 在项目中, 不要轻易使用 panic
		// panic 会中断整个项目, 并且抛出这个 panic
		// 只有在确实无法正常完成业务逻辑的时候, 才使用 panic
		panic(err.Error())
	}
	return dbConn
}

func AddUserCredential(loginName string, pwd string) error {
	// openConn 的时候实际上并不会连接数据库
	// 只有在 Prepare 的时候才会真正连接数据库
	db := openConn()
}

func GetUserCredential(loginName string) (string, error) {
	db := openConn()
}
