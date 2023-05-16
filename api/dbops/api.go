package dbops

import (
	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	// openConn 的时候实际上并不会连接数据库
	// 只有在 Prepare 的时候才会真正连接数据库

}

func GetUserCredential(loginName string) (string, error) {

}
