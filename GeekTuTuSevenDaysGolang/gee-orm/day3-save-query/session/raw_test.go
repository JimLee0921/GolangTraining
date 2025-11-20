package session

import (
	"database/sql"
	"day1-database-sql/dialect"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite")
)

// m *testing.M 是整个测试框架的入口控制器
// 允许在所有测试开始之前做初始化，在所有测试结束后做清理，m.Run() 会执行所有 TestXXX
func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite", "../gee.db")
	code := m.Run() // 所有 TestXXX 在这里执行
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text)").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "JimLee", "JamesBond").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2, but got", count)
	}
}

func TestSession_QueryRow(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	row := s.Raw("SELECT count(*) FROM User;").QueryRow()
	var count int
	if err := row.Scan(&count); err != nil || count != 0 {
		t.Fatal("failed to query db", err)
	}
}
