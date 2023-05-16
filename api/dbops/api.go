package dbops

import (
	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {

}

func GetUserCredential(loginName string) (string, error) {

}
