package dbops

import "testing"

// 说下测试顺序：
// init(dblogin, truncate tables) -> run tests -> clear data(truncate tables)
// 为什么要在 init 中 truncate tables 呢？
// 因为我们的测试是在真实的数据库上进行的，所以我们要保证测试的数据是干净的，不然会影响测试结果

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("test_user", "123456")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("test_user")
	if pwd != "123456" || err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUserCredential("test_user", "123456")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("test_user")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test_user failed")
	}
}
