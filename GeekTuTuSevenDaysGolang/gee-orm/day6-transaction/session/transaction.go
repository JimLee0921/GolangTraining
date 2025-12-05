package session

import (
	"day1-database-sql/mylog"
)

// 封装事务接口

// Begin 开启事务，获得一个 *sql.Tx 存入 s.tx
func (s *Session) Begin() (err error) {
	mylog.Info("transaction begin")
	if s.tx, err = s.db.Begin(); err != nil {
		mylog.Error(err)
		return
	}
	return
}

// Commit 提交事务，调用 s.tx.Commit()
func (s *Session) Commit() (err error) {
	mylog.Info("transaction commit")
	if err = s.tx.Commit(); err != nil {
		mylog.Error(err)
	}
	return
}

// Rollback 回滚事务，调用 s.tx.Rollback()
func (s *Session) Rollback() (err error) {
	mylog.Info("transaction rollback")
	if err = s.tx.Rollback(); err != nil {
		mylog.Error(err)
	}
	return
}
