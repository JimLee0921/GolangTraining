package geeorm

import (
	"day1-database-sql/session"
	"errors"
	"reflect"
	"testing"
)
import _ "modernc.org/sqlite"

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	engine, err := NewEngine("sqlite", "gee.db")
	if err != nil {
		t.Fatal("failed to connect db", err)
	}
	return engine
}

func TestNewEngine(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
}

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func transactionRollback(t *testing.T) {
	engin := OpenDB(t)
	defer engin.Close()
	s := engin.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engin.Transaction(func(session *session.Session) (result any, err error) {
		_ = session.Model(&User{}).CreateTable()
		_, err = session.Insert(&User{"Tom", 18})
		return nil, errors.New("error")
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}

func transactionCommit(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := engine.Transaction(func(s *session.Session) (result any, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit", u)
	}
}

func TestEngine_Migrate(t *testing.T) {
	engine := OpenDB(t)
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text PRIMARY KEY, XXX integer);").Exec()
	_, _ = s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	engine.Migrate(&User{})
	rows, _ := s.Raw("SELECT * FROM User").Query()
	columns, _ := rows.Columns()
	if !reflect.DeepEqual(columns, []string{"Name", "Age"}) {
		t.Fatal("failed to migrate table user, got columns", columns)
	}
}
